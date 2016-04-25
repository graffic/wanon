package quotes

import (
	"fmt"

	"github.com/graffic/wanon/bot"
)

type randomQuote struct {
	quoteHandler
}

func (handler *randomQuote) Check(message *bot.Message) int {
	return handler.check("/rquote", message)
}

func (handler *randomQuote) Handle(message *bot.Message) {
	quote, err := handler.storage.RQuote(message.Chat.ID)
	if err != nil {
		log.Error(fmt.Sprint(err))
		return
	}

	if quote == nil {
		_, err = message.Reply("I'm empty! Add quotes to me")
	} else {
		_, err = message.Send(fmt.Sprintf("<%s> %s", quote.SaidBy, quote.What))
	}

	if err != nil {
		log.Error(fmt.Sprint(err))
	}
}

// CreateRandomQuote does nothing
func CreateRandomQuote(context *bot.Context) bot.Handler {
	return &randomQuote{createQuoteHandler(context)}
}
