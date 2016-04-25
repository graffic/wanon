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

func (handler *listHandler) Check(message *bot.Message) int {
	return handler.check("/list", message)
}

func (handler *listHandler) Handle(message *bot.Message) {
	items := strings.Split(message.Text, " ")
	amount := len(items)
	skip := 0

	if amount < 2 {
		message.Reply("You forgot the chat id")
		return
	}
	if amount == 3 {
		skip, _ = strconv.Atoi(items[2])
	}

	log.Debug("List quotes on %d skip: %d", items[1], skip)
	quotes, _ := handler.storage.List(items[1], skip)
	var result string
	for _, quote := range *quotes {
		quoteStr := fmt.Sprintf("%s: <%s> %s\n", quote.ID.Hex(), quote.SaidBy, quote.What)
		result += quoteStr
	}
	log.Debug(result)
	message.Reply(result)

}
