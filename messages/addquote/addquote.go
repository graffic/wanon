package addquote

import (
	"strings"

	"github.com/graffic/wanon/bot"
	"github.com/graffic/wanon/telegram"
)

type addQuote struct{}

func (handler *addQuote) Check(message *telegram.Message) int {
	isAddQuote := strings.Index(message.Text, "/addquote") == 0
	isReply := message.ReplyToMessage != nil

	if isAddQuote && isReply {
		return bot.RouteAccept
	}
	return bot.RouteNothing
}

func (handler *addQuote) Handle(message *telegram.Message) {

}

// Create does nothing
func Create() bot.Handler {
	return new(addQuote)
}
