package quotes

import (
	"fmt"

	"github.com/graffic/wanon/bot"
	"github.com/op/go-logging"
)

type addQuote struct {
	quoteHandler
}

var log = logging.MustGetLogger("wanon.messages.quotes")

func (handler *addQuote) Check(message *bot.Message) int {
	res := handler.check("/addquote", message)

	if res == bot.RouteAccept && message.ReplyToMessage == nil {
		message.Reply("To add a quote use /addquote in a reply")
		res = bot.RouteStop
	}

	return res
}

func (handler *addQuote) Handle(message *bot.Message) {
	quote := Quote{
		AddedBy: message.From.Username,
		SaidBy:  message.ReplyToMessage.From.Username,
		When:    message.Date,
		What:    message.ReplyToMessage.Text,
	}

	err := handler.storage.AddQuote(message.Chat.ID, &quote)
	if err != nil {
		log.Fatal(err)
	}

	quoteRendered := fmt.Sprintf("<%s> %s", quote.SaidBy, quote.What)
	log.Info("Quote Added: " + quoteRendered)
	message.Send(quoteRendered + "\nprocesado correctamente, siguienteeeeeee!!!!")
}

// CreateAddQuote does nothing
func CreateAddQuote(context *bot.Context) bot.Handler {
	return &addQuote{createQuoteHandler(context)}
}
