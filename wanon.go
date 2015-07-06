package main

import (
	"fmt"
	"github.com/graffic/wanon/conf"
	"github.com/graffic/wanon/telegram"
	"log"
	"os"
)

func printMessages(messages chan *telegram.TelegramMessage) {
	for {
		message := <-messages
		fmt.Println(message.Text)
	}
}

func main() {
	conf, err := conf.Load("conf.yaml")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	api := telegram.NewTelegramApi(conf.Token)
	result := api.GetMe()
	fmt.Println(result.Username, "online")

	channel := make(chan *telegram.TelegramMessage)
	go telegram.ProcessUpdates(&api, channel)
	printMessages(channel)
}
