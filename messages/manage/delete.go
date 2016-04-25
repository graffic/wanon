package manage

import (
	"fmt"
	"strings"

	"github.com/graffic/wanon/bot"
)

type quoteDeleter interface {
	Delete(chat string, oid string) error
}
type deleteHandler struct {
	*manageHandler
	storage quoteDeleter
}

func (handler *deleteHandler) Check(message *bot.Message) int {
	return handler.check("/delete", message)
}

func (handler *deleteHandler) Handle(message *bot.Message) {
	items := strings.Split(message.Text, " ")

	if len(items) != 3 {
		return
	}

	err := handler.storage.Delete(items[1], items[2])
	if err != nil {
		message.Reply(fmt.Sprint(err))
	} else {
		message.Reply("Deleted")
	}
}
