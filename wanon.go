package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/graffic/wanon/bot"
	"github.com/graffic/wanon/messages/ignorechat"
	"github.com/graffic/wanon/messages/manage"
	"github.com/graffic/wanon/messages/quotes"
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
	logging.SetLevel(logging.INFO, "wanon.telegram")
}

// checkFatal exits on fatal
func checkFatal(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func sigQuit() {
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGQUIT)
		buf := make([]byte, 1<<20)
		for {
			<-sigs
			runtime.Stack(buf, true)
			log.Warning("=== received SIGQUIT ===\n*** goroutine dump...\n%s\n*** end\n", buf)
		}
	}()
}

func main() {
	initLogging()
	log.Debug("Wanon booting")
	sigQuit()

	context, err := bot.CreateContext("conf.yaml")
	checkFatal(err)

	log.Info("All systems nominal")

	channel := make(chan *telegram.Message)
	go context.API.ProcessUpdates(channel)

	router := bot.Router{}
	router.AddHandler(ignorechat.Create(context.Conf))
	router.AddHandler(quotes.CreateAddQuote())
	router.AddHandler(quotes.CreateRandomQuote())
	router.AddHandler(manage.Create(context.Conf))
	router.RouteMessages(channel, context)
}
