// Package manage move quotes from one chat to another
package manage

import (
	"fmt"

	"github.com/graffic/wanon/bot"
)

type quoteMover interface {
	Move(from string, to string) (int, error)
}

type moveHandler struct {
	*manageHandler
	storage quoteMover
}

func (handler *moveHandler) Handle(context *bot.MessageContext) {
	from := context.Params["from"]
	to := context.Params["to"]
	amount, err := handler.storage.Move(from, to)

	if err != nil {
		logger.Errorf("%v", err)
		return
	}

	context.Message.Reply(fmt.Sprintf("Moved %d quotes", amount))
}
