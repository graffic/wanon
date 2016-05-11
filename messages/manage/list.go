package manage

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/graffic/wanon/bot"
	"github.com/graffic/wanon/messages/quotes"
)

type quoteLister interface {
	List(chat string, amountToSkip int) (*[]quotes.Quote, error)
}

type listHandler struct {
	*manageHandler
	storage quoteLister
}

func (handler *listHandler) Handle(context *bot.MessageContext) {
	skip := 0
	chat := context.Params["chat"]

	if strings.Index(chat, "quotes_") != 0 {
		context.Message.Reply("Wrong collection")
		return
	}

	givenSkip, ok := context.Params["skip"]
	if ok {
		skip, _ = strconv.Atoi(givenSkip)
	}

	logger.Debug("List quotes on %d skip: %d", chat, skip)
	quotes, _ := handler.storage.List(chat, skip)
	var result string
	for _, quote := range *quotes {
		quoteStr := fmt.Sprintf("%s: <%s> %s\n", quote.ID.Hex(), quote.SaidBy, quote.What)
		result += quoteStr
	}
	logger.Debug(result)
	context.Message.Reply(result)

}
