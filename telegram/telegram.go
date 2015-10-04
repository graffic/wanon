package telegram

import (
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

func call(url string) (Response, error) {
	var response Response
	// log.Debug(url)
	resp, err := http.Get(url)

	if err != nil {
		return response, err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	log.Debug(string(bytes))
	json.Unmarshal(bytes, &response)

	return response, nil
}

// GetMe checks the token validity
func (api *API) GetMe() (User, error) {
	var user User
	response, err := call(api.baseURL + "getMe")

	if err != nil {
		return user, err
	}

	json.Unmarshal(response.Result, &user)
	return user, nil
}

// GetUpdates from telegram
func (api *API) GetUpdates(offset int) []Update {
	var update []Update
	url := fmt.Sprintf("%sgetUpdates?offset=%d&timeout=5", api.baseURL, offset)
	log.Info("getting updates...")
	response, err := call(url)
	if err != nil {
		fmt.Println(err)
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
