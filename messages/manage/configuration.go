package manage

import (
	"github.com/graffic/wanon/bot"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("wanon.messages.manage")

type manageHandler struct {
	allowed map[int]bool
}

type configuration struct {
	Admins []int
}

func (handler *manageHandler) Check(context *bot.MessageContext) int {
	id := context.Message.Chat.ID

	if handler.allowed[id] {
		return bot.RouteAccept
	}
	logger.Debug("Not allowed: %d", id)
	return bot.RouteStop
}

func createManageHandler(context *bot.BotContext) *manageHandler {
	myConf := new(configuration)
	context.Conf.Get(myConf)
	logger.Notice("Manage only from: %v", myConf.Admins)

	allowed := make(map[int]bool)
	for _, chatID := range myConf.Admins {
		allowed[chatID] = true
	}

	return &manageHandler{allowed}
}

// Handlers stores handlers
type Handlers interface {
	AddHandler(definition string, givenHandler bot.Handler)
}

// Setup the manage commands
func Setup(handlers Handlers, context *bot.BotContext) {
	handler := createManageHandler(context)
	storage := &manageStorage{context.Storage}

	list := &listHandler{handler, storage}
	handlers.AddHandler("/list :chat", list)
	handlers.AddHandler("/list :chat :skip", list)

	handlers.AddHandler("/chats", &chatsHandler{handler, storage})
	handlers.AddHandler("/delete :chat :message", &deleteHandler{handler, storage})
	handlers.AddHandler("/move :from :to", &moveHandler{handler, storage})
}
