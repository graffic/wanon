package quotes

import (
	"fmt"

	"github.com/graffic/wanon/bot"
)

type addQuote struct {
	quoteHandler
}

func (handler *addQuote) Check(context *bot.MessageContext) int {
	message := context.Message
	res := handler.quoteHandler.Check(context)

	if res == bot.RouteAccept && message.ReplyToMessage == nil {
		message.Reply("To add a quote use /addquote in a reply")
		res = bot.RouteStop
	}

	return res
}

func (handler *addQuote) Handle(context *bot.MessageContext) {
	message := context.Message
	quote := Quote{
		AddedBy: message.From.Username,
		SaidBy:  message.ReplyToMessage.From.Username,
		When:    message.Date,
		What:    message.ReplyToMessage.Text,
	}

	err := handler.storage.AddQuote(message.Chat.ID, &quote)
	if err != nil {
		logger.Fatal(err)
	}

	quoteRendered := fmt.Sprintf("<%s> %s", quote.SaidBy, quote.What)
	logger.Infof("Quote Added: %s", quoteRendered)
	message.Send(quoteRendered + "\nprocesado correctamente, siguienteeeeeee!!!!")
}

// CreateAddQuote does nothing
func CreateAddQuote(context *bot.BotContext) bot.Handler {
	return &addQuote{createQuoteHandler(context)}
}
