package telegram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TelegramResponse struct {
	Ok     bool            `json:ok`
	Result json.RawMessage `json:result`
}

type TelegramUser struct {
	Id        int    `json:id`
	FirstName string `json:first_name`
	LastName  string `json:last_name`
	Username  string `json:username`
}

type TelegramApi struct {
	baseUrl string
}

type TelegramUpdate struct {
	UpdateId int             `json:"update_id"`
	Message  TelegramMessage `json:message`
}

type TelegramMessage struct {
	MessageId int
	Text      string
}

func call(url string) (TelegramResponse, error) {
	var response TelegramResponse
	//fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		return response, err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bytes))
	json.Unmarshal(bytes, &response)

	return response, nil
}

func (api *TelegramApi) GetMe() TelegramUser {
	var user TelegramUser
	response, err := call(api.baseUrl + "getMe")

	if err != nil {
		fmt.Println(err)
		return user
	}

	json.Unmarshal(response.Result, &user)
	return user
}

func (api *TelegramApi) GetUpdates(offset int) []TelegramUpdate {
	var update []TelegramUpdate
	url := fmt.Sprintf("%sgetUpdates?offset=%d&timeout=5", api.baseUrl, offset)
	response, err := call(url)
	if err != nil {
		fmt.Println(err)
		return update
	}
	json.Unmarshal(response.Result, &update)

	return update
}

func NewTelegramApi(token string) TelegramApi {
	baseUrl := "https://api.telegram.org/bot" + token + "/"
	return TelegramApi{baseUrl}
}

func ProcessUpdates(api *TelegramApi, out chan *TelegramMessage) {
	lastMessageId := 0
	for 1 == 1 {
		update := api.GetUpdates(lastMessageId)

		for _, element := range update {
			if element.UpdateId >= lastMessageId {
				lastMessageId = element.UpdateId + 1
			}
			out <- &element.Message
		}
	}
}
