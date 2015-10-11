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
		l.Error(fmt.Sprint(err))
		return
	}

	answer := telegram.AnswerBack{API: context.API, Message: message}

	if quote == nil {
		_, err = answer.Reply("I'm empty! Add quotes to me")
	} else {
		_, err = answer.Send(fmt.Sprintf("<%s> %s", quote.SaidBy, quote.What))
	}

	if err != nil {
		l.Error(fmt.Sprint(err))
	}
}

// CreateRandomQuote does nothing
func CreateRandomQuote() bot.Handler {
	return new(randomQuote)
}
