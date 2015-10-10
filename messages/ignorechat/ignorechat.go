package ignorechat

import (
	"github.com/graffic/wanon/bot"
	"github.com/graffic/wanon/telegram"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("wanon.messages.ignorechat")

type ignoreHandler struct {
	allowed map[int]bool
}

// IgnoreConfiguration configuration for the ignore message
type IgnoreConfiguration struct {
	Allow []int
}

func (handler *ignoreHandler) Check(message *telegram.Message, context *bot.Context) int {
	if handler.allowed[message.Chat.ID] {
		log.Debug("Accepted")
		return bot.RouteNothing
	}
	log.Debug("Ignored")
	return bot.RouteStop
}

func (handler *ignoreHandler) Handle(message *telegram.Message, context *bot.Context) {}

// Create creates the ignore handler
func Create(conf *bot.ConfService) bot.Handler {
	myConf := new(IgnoreConfiguration)
	conf.Get(myConf)
	log.Notice("Allowing only from: %d", myConf.Allow)

	allowed := make(map[int]bool)
	for _, chatID := range myConf.Allow {
		allowed[chatID] = true
	}

	return &ignoreHandler{allowed}
}
