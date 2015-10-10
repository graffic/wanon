package quotes

import (
	"fmt"
	"strings"

	"github.com/graffic/wanon/bot"
	"github.com/graffic/wanon/telegram"
)

type randomQuote struct{}

func (handler *randomQuote) Check(message *telegram.Message, context *bot.Context) int {
	isRandom := strings.Index(message.Text, "/rquote") == 0
	isNotReply := message.ReplyToMessage == nil

	if isRandom && isNotReply {
		l.Debug("Random requested")
		return bot.RouteAccept
	}
	return bot.RouteNothing
}

func (handler *randomQuote) Handle(message *telegram.Message, context *bot.Context) {
	quotes := quoteStorage{context.Storage}
	quote, err := quotes.RQuote(message.Chat.ID)
	if err != nil {
		l.Debug("%s", err)
		return
	}
	if quote == nil {
		return
	}

	_, err = context.API.SendMessage(&telegram.SendMessage{
		ChatID: message.Chat.ID,
		Text:   fmt.Sprintf("<%s> %s", quote.SaidBy, quote.What),
	})
	if err != nil {
		l.Error(fmt.Sprint(err))
	}
}

// CreateRandomQuote does nothing
func CreateRandomQuote() bot.Handler {
	return new(randomQuote)
}
