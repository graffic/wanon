package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/graffic/wanon/bot"
	"github.com/graffic/wanon/messages"
	"github.com/graffic/wanon/migrations"
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

// checkFatal treats errors as fatal errors (log.Fatal exits)
func checkFatal(err error) {
	if err != nil {
		log.Fatal(err)
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

	err = migrations.Run(context.Storage)
	checkFatal(err)

	log.Info("All systems nominal")

	channel := make(chan *telegram.Message)
	go context.API.ProcessUpdates(channel)

	router := &bot.Routes{}
	messages.Setup(router, context)
	bot.MainLoop(channel, router)
}
