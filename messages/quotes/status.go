package quotes

import (
	"fmt"

	"github.com/graffic/wanon/bot"
)

type quoteStatus struct {
	quoteHandler
}

func (handler *quoteStatus) Handle(context *bot.MessageContext) {
	message := context.Message
	amount := 0

	collName := fmt.Sprintf("quotes_%d", message.Chat.ID)
	col, err := handler.storage.ejdb.GetColl(collName)
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
