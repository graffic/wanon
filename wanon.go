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
		"%{color}%{time:15:04:05.000} %{level:5.5s} %{module} >%{color:reset} %{message}",
	)
	formatter := logging.NewBackendFormatter(backend, format)

	// Set the backends to be used.
	logging.SetBackend(formatter)
}

func checkFatal(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func createAPI(conf *bot.ConfService) (*telegram.API, error) {
	var apiConf telegram.Configuration
	conf.Get(&apiConf)
	api := telegram.NewAPI(apiConf)
	result, err := api.GetMe()
	if err != nil {
		return nil, err
	}
	log.Info("%s online", result.Username)

	return &api, nil
}

func main() {
	initLogging()
	log.Debug("Wanon booting")

	conf, err := bot.LoadConf("conf.yaml")
	checkFatal(err)

	_, err2 := bot.NewStorage(conf)
	checkFatal(err2)

	api, err3 := createAPI(conf)
	checkFatal(err3)

	log.Info("All systems nominal")

	channel := make(chan *telegram.Message)
	go api.ProcessUpdates(channel)
	router := bot.Router{}
	router.AddHandler(messages.CreateIgnore(conf))
	router.RouteMessages(channel)
}
