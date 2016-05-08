package quotes

import (
	"fmt"
	"strconv"

	"github.com/graffic/wanon/bot"
)

type quoteStatus struct {
	quoteHandler
}

func (handler *quoteStatus) Handle(context *bot.MessageContext) {
	message := context.Message
	amount := 0

	col, err := handler.storage.GetColl(strconv.Itoa(message.Chat.ID))
	if err != nil && err.ErrorCode != 9000 {
		logger.Errorf("%v", err)
		return
	}

	if col != nil {
		amount, err = col.Count("{}")
		if err != nil {
			logger.Errorf("%v", err)
		}
	}

	message.Reply(fmt.Sprintf("Tengo %d quotes para este chat", amount))
}
