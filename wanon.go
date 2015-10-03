package main

import (
	"os"

	"github.com/graffic/wanon/bot"
	"github.com/graffic/wanon/messages"
	"github.com/graffic/wanon/telegram"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("wanon")

func initLogging() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	format := logging.MustStringFormatter(
		"%{color}%{time:15:04:05.000} %{module}💾 %{level:.5s} %{color:reset} %{message}",
	)
	formatter := logging.NewBackendFormatter(backend, format)

	// Set the backends to be used.
	logging.SetBackend(formatter)
}

func createAPI(conf *bot.ConfService) *telegram.API {
	var apiConf telegram.Configuration
	conf.Get(&apiConf)
	api := telegram.NewAPI(apiConf)
	result := api.GetMe()
	log.Info("%s online", result.Username)

	return &api
}

func main() {
	initLogging()

	conf, err := bot.LoadConf("conf.yaml")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	api := createAPI(conf)

	channel := make(chan *telegram.Message)
	go api.ProcessUpdates(channel)
	router := bot.Router{}
	router.AddHandler(messages.CreateIgnore(conf))
	router.RouteMessages(channel)
}
