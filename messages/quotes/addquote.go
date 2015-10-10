package quotes

import (
	"strings"

	"github.com/graffic/wanon/bot"
	"github.com/graffic/wanon/telegram"
	"github.com/op/go-logging"
)

type addQuote struct{}

var l = logging.MustGetLogger("wanon.messages.quotes")

func (handler *addQuote) Check(message *telegram.Message, context *bot.Context) int {
	isAddQuote := strings.Index(message.Text, "/addquote") == 0
	isReply := message.ReplyToMessage != nil

	if isAddQuote && isReply {
		l.Info("Adding quote")
		return bot.RouteAccept
	}
	return bot.RouteNothing
}

func (handler *addQuote) Handle(message *telegram.Message, context *bot.Context) {
	quotes := quoteStorage{context.Storage}
	quote := Quote{
		AddedBy: message.From.Username,
		SaidBy:  message.ReplyToMessage.From.Username,
		When:    message.Date,
		What:    message.ReplyToMessage.Text,
	}

	err := quotes.AddQuote(message.Chat.ID, &quote)
	if err != nil {
		l.Fatal(err)
	}
	l.Info("Quote Added: <%s> %s", quote.SaidBy, quote.What)
}

// CreateAddQuote does nothing
func CreateAddQuote() bot.Handler {
	return new(addQuote)
}
