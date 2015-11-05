package telegram

import (
	"encoding/json"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("wanon.telegram")

// apiImpl holds the API configuration and methods
type apiImpl struct {
	Request
}

// Configuration settings
type Configuration struct {
	Token string
}

// API holds the API methods
type API interface {
	GetMe() (User, error)
	GetUpdates(int) []Update
	ProcessUpdates(chan *Message)
	SendMessage(*SendMessage) (*Message, error)
}

// GetMe checks the token validity
func (api *apiImpl) GetMe() (User, error) {
	var user User
	response, err := api.Call("getMe", nil)

	if err != nil {
		return user, err
	}

	json.Unmarshal(response.Result, &user)
	return user, nil
}

// GetUpdates from telegram
func (api *apiImpl) GetUpdates(offset int) []Update {
	var update []Update

	log.Debug("getting updates...")
	response, err := api.Call("getUpdates", GetUpdates{offset, 5})
	if err != nil {
		return update
	}
	json.Unmarshal(response.Result, &update)

	return update
}

// NewAPI creates a new api from a token
func NewAPI(httpClient HTTPClient, conf Configuration) API {
	baseURL := "https://api.telegram.org/bot" + conf.Token + "/"
	return &apiImpl{NewRequest(httpClient, baseURL)}
}

// ProcessUpdates from the telegram api
// Note: I don't like this here somehow.
func (api *apiImpl) ProcessUpdates(out chan *Message) {
	lastMessageID := 0
	for {
		update := api.GetUpdates(lastMessageID)

		for _, element := range update {
			if element.UpdateID >= lastMessageID {
				lastMessageID = element.UpdateID + 1
			}
			out <- &element.Message
		}
	}
}

// SendMessage action
func (api *apiImpl) SendMessage(message *SendMessage) (*Message, error) {
	response, err := api.Call("sendMessage", message)
	if err != nil {
		return nil, err
	}
	var msg Message

	json.Unmarshal(response.Result, &msg)
	return &msg, nil
}
