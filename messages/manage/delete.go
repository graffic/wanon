package manage

import (
	"fmt"

	"github.com/graffic/wanon/bot"
)

type quoteDeleter interface {
	Delete(chat string, oid string) error
}
type deleteHandler struct {
	*manageHandler
	storage quoteDeleter
}

func (handler *deleteHandler) Handle(context *bot.MessageContext) {
	err := handler.storage.Delete(context.Params["chat"], context.Params["message"])
	if err != nil {
		context.Message.Reply(fmt.Sprint(err))
	} else {
		context.Message.Reply("Deleted")
	}
}
