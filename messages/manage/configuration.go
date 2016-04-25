package manage

import (
	"strings"

	"github.com/graffic/wanon/bot"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("wanon.messages.manage")

type manageHandler struct {
	allowed map[int]bool
}

type configuration struct {
	Admins []int
}

func (handler *manageHandler) check(command string, message *bot.Message) int {
	isCommand := strings.Index(message.Text, command) == 0

	if handler.allowed[message.Chat.ID] && isCommand {
		return bot.RouteAccept
	}
	if isCommand {
		log.Debug("Not allowed: %d", message.Chat.ID)
		return bot.RouteStop
	}
	return bot.RouteNothing
}

func createManageHandler(context *bot.Context) *manageHandler {
	myConf := new(configuration)
	context.Conf.Get(myConf)
	log.Notice("Manage only from: %v", myConf.Admins)

	allowed := make(map[int]bool)
	for _, chatID := range myConf.Admins {
		allowed[chatID] = true
	}

	return &manageHandler{allowed}
}

// Setup the manage commands
func Setup(router *bot.Router, context *bot.Context) {
	handler := createManageHandler(context)
	storage := &manageStorage{context.Storage}

	router.AddHandler(&listHandler{handler, storage})
	router.AddHandler(&chatsHandler{handler, storage})
	router.AddHandler(&deleteHandler{handler, storage})
	router.AddHandler(&moveHandler{handler, storage})
}
