package quotes

import (
	"fmt"
	"strconv"

	"github.com/graffic/wanon/bot"
)

type quoteStatus struct {
	quoteHandler
}

func (handler *quoteStatus) Check(message *bot.Message) int {
	return handler.check("/status", message)
}

func (handler *quoteStatus) Handle(message *bot.Message) {
	amount := 0

	col, err := handler.storage.GetColl(strconv.Itoa(message.Chat.ID))
	if err != nil && err.ErrorCode != 9000 {
		log.Error("%v", err)
		return
	}

	if col != nil {
		amount, err = col.Count("{}")
		if err != nil {
			log.Error("%v", err)
		}
	}

	message.Reply(fmt.Sprintf("Tengo %d quotes para este chat", amount))
}
