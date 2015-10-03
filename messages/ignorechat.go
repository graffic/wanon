package messages

import (
	"fmt"

	"github.com/graffic/wanon/bot"
	"github.com/graffic/wanon/telegram"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("wanon.messages")

// IgnoreConfiguration configuration for the ignore message
type IgnoreConfiguration struct {
	Allow []int
}

func createIgnoreCheck(conf *IgnoreConfiguration) bot.HandlerCheck {
	log.Notice("Allowing only from: %d", conf.Allow)
	allowed := make(map[int]bool)
	for _, chatID := range conf.Allow {
		allowed[chatID] = true
	}

	return func(message *telegram.Message) int {
		if allowed[message.Chat.ID] {
			fmt.Println("Accepted")
			return bot.RouteNothing
		}
		fmt.Println("Ignored")
		return bot.RouteStop
	}
}

func ignoreHandle(message *telegram.Message) {}

// CreateIgnore creates the ignore handler
func CreateIgnore(conf *bot.ConfService) *bot.Handler {
	myConf := new(IgnoreConfiguration)
	conf.Get(myConf)

	handler := bot.NewHandler(createIgnoreCheck(myConf), ignoreHandle)
	return &handler
}
