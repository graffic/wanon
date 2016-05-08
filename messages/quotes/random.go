package quotes

import (
	"fmt"

	"github.com/graffic/wanon/bot"
)

type randomQuote struct {
	quoteHandler
}

func (handler *randomQuote) Handle(context *bot.MessageContext) {
	message := context.Message
	quote, err := handler.storage.RQuote(message.Chat.ID)
	if err != nil {
		logger.Error(fmt.Sprint(err))
		return
	}

	if quote == nil {
		_, err = message.Reply("I'm empty! Add quotes to me")
	} else {
		_, err = message.Send(fmt.Sprintf("<%s> %s", quote.SaidBy, quote.What))
	}

	if err != nil {
		logger.Error(fmt.Sprint(err))
	}
}

// CreateRandomQuote does nothing
func CreateRandomQuote(context *bot.BotContext) bot.Handler {
	return &randomQuote{createQuoteHandler(context)}
}
