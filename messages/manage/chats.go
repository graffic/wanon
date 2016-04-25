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

func (handler *chatsHandler) Check(message *bot.Message) int {
	return handler.check("/chats", message)
}

func (handler *chatsHandler) Handle(message *bot.Message) {
	log.Debug("Chat database status")
	chats, err := handler.storage.Chats()
	if err != nil {
		log.Error(fmt.Sprint(err))
		return
	}

	var result string
	for _, chat := range *chats {
		result += fmt.Sprintf("%s: %d\n", chat.Name, chat.Records)
	}

	message.Reply(result)
}
