package messages

import (
	"github.com/graffic/wanon/bot"
	"github.com/graffic/wanon/messages/manage"
	"github.com/graffic/wanon/messages/quotes"
)

// Handlers stores handlers
type Handlers interface {
	AddHandler(definition string, givenHandler bot.Handler)
}

// Setup bot messages
func Setup(handlers Handlers, context *bot.BotContext) {
	quotes.Setup(handlers, context)
	manage.Setup(handlers, context)
}
