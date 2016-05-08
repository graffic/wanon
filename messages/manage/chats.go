package manage

import (
	"fmt"

	"github.com/graffic/wanon/bot"
)

type chatLister interface {
	Chats() (*[]collectionMetadata, error)
}

type chatsHandler struct {
	*manageHandler
	storage chatLister
}

func (handler *chatsHandler) Handle(context *bot.MessageContext) {
	logger.Debug("Chat database status")
	chats, err := handler.storage.Chats()
	if err != nil {
		logger.Error(fmt.Sprint(err))
		return
	}

	var result string
	for _, chat := range *chats {
		result += fmt.Sprintf("%s: %d\n", chat.Name, chat.Records)
	}

	context.Message.Reply(result)
}
