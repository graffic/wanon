package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("wanon.telegram")

// API holds the API configuration and methods
type API struct {
	baseURL string
}

// Configuration settings
type Configuration struct {
	Token string
}

func (api *API) call(method string, in interface{}) (*Response, error) {
	url := fmt.Sprintf("%s%s", api.baseURL, method)
	var response *http.Response
	var err error

	if in == nil {
		response, err = http.Get(url)
	} else {
		outData, err := json.Marshal(in)
		if err != nil {
			return nil, err
		}
		log.Debug("Request: " + string(outData))
		response, err = http.Post(url, "application/json", bytes.NewBuffer(outData))
	}
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	log.Debug("Response: " + string(bytes))

	var out Response
	json.Unmarshal(bytes, &out)

	return &out, nil
}

// GetMe checks the token validity
func (api *API) GetMe() (User, error) {
	var user User
	response, err := api.call("getMe", nil)

	if err != nil {
		return user, err
	}

	json.Unmarshal(response.Result, &user)
	return user, nil
}

// GetUpdates from telegram
func (api *API) GetUpdates(offset int) []Update {
	var update []Update

	log.Info("getting updates...")
	response, err := api.call("getUpdates", GetUpdates{offset, 5})
	if err != nil {

		return update
	}
	json.Unmarshal(response.Result, &update)

	return update
}

// NewAPI creates a new api from a token
func NewAPI(conf Configuration) API {
	baseURL := "https://api.telegram.org/bot" + conf.Token + "/"
	return API{baseURL}
}

// ProcessUpdates from the telegram api
// Note: I don't like this here somehow.
func (api *API) ProcessUpdates(out chan *Message) {
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
func (api *API) SendMessage(message *SendMessage) (*Message, error) {
	response, err := api.call("sendMessage", message)
	if err != nil {
		return nil, err
	}
	var msg Message

	json.Unmarshal(response.Result, &msg)
	return &msg, nil
}
