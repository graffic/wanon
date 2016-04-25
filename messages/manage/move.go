// Package manage move quotes from one chat to another
package manage

import (
	"fmt"
	"strings"

	"github.com/graffic/wanon/bot"
)

type quoteMover interface {
	Move(from string, to string) (int, error)
}

type moveHandler struct {
	*manageHandler
	storage quoteMover
}

func (handler *moveHandler) Check(message *bot.Message) int {
	return handler.check("/move", message)
}

func (handler *moveHandler) Handle(message *bot.Message) {
	items := strings.Split(message.Text, " ")
	if len(items) != 3 {
		return
	}

	amount, err := handler.storage.Move(items[1], items[2])

	if err != nil {
		log.Error("%v", err)
		return
	}

	message.Reply(fmt.Sprintf("Moved %d quotes", amount))
}
