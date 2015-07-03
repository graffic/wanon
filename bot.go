package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Configuration struct {
	Token string
}

func loadConfiguration(fileName string) (Configuration, error) {
	var conf Configuration

	file, err := os.Open(fileName)
	if err != nil {
		return conf, err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

type TelegramResponse struct {
	Ok     bool            `json:ok`
	Result json.RawMessage `json:result`
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

func (api *TelegramApi) getMe() TelegramUser {
	var user TelegramUser
	response, err := call(api.baseUrl + "getMe")

	if err != nil {
		fmt.Println(err)
		return user
	}

	json.Unmarshal(response.Result, &user)
	return user
}

func (api *TelegramApi) getUpdates(offset int) []TelegramUpdate {
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

func buildTelegramApi(token string) TelegramApi {
	baseUrl := "https://api.telegram.org/bot" + token + "/"
	return TelegramApi{baseUrl}
}

func retrieveMessages(api *TelegramApi) {
	lastMessageId := 0
	for 1 == 1 {
		update := api.getUpdates(lastMessageId)

		for _, element := range update {
			if element.UpdateId >= lastMessageId {
				lastMessageId = element.UpdateId + 1
			}
			fmt.Println(element.Message.Text)
		}
		time.Sleep(time.Second)
	}
}

func main() {
	conf, err := loadConfiguration("conf.yaml")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	api := buildTelegramApi(conf.Token)
	result := api.getMe()
	fmt.Println(result.Username, "online")
	retrieveMessages(&api)
}
