package manage

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/graffic/wanon/bot"
	"github.com/graffic/wanon/telegram"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("wanon.messages.manage")

type handler struct {
	allowed map[int]bool
}

type configuration struct {
	Admins []int
}

type manager struct {
	storage managerStorage
	answer  telegram.AnswerBack
}

func (handler *handler) Check(message *telegram.Message, context *bot.Context) int {
	isManage := strings.Index(message.Text, "/manage") == 0
	if handler.allowed[message.Chat.ID] && isManage {
		log.Debug("Accepted")
		return bot.RouteAccept
	}
	if isManage {
		log.Debug(fmt.Sprintf("Not allowed: %d", message.Chat.ID))
	}
	return bot.RouteStop
}

func (handler *handler) Handle(message *telegram.Message, context *bot.Context) {
	actions := manager{
		managerStorage{context.Storage},
		telegram.AnswerBack{API: context.API, Message: message},
	}

	switch {
	case strings.Index(message.Text, "/manage list") == 0:
		actions.list(message.Text)
	case strings.Index(message.Text, "/manage delete") == 0:
		actions.delete(message.Text)
	case strings.Index(message.Text, "/manage chats") == 0:
		actions.chats()
	default:
		help()
	}
}

func (actions *manager) list(text string) {
	items := strings.Split(text, " ")
	amount := len(items)
	skip := 0

	if amount < 3 {
		return
	}
	if amount == 4 {
		skip, _ = strconv.Atoi(items[3])
	}

	quotes, _ := actions.storage.List(items[2], skip)
	var result string
	for _, quote := range *quotes {
		quoteStr := fmt.Sprintf("%s: <%s> %s\n", quote.ID.Hex(), quote.SaidBy, quote.What)
		result += quoteStr
	}
	actions.answer.Reply(result)
}

func (actions *manager) chats() {
	log.Debug("Chat database status")
	chats, err := actions.storage.Chats()
	if err != nil {
		log.Error(fmt.Sprint(err))
		return
	}

	var result string
	for _, chat := range *chats {
		result += fmt.Sprintf("%s: %d\n", chat.Name, chat.Records)
	}

	actions.answer.Reply(result)
}

func (actions *manager) delete(text string) {
	items := strings.Split(text, " ")

	if len(items) != 4 {
		return
	}

	err := actions.storage.Delete(items[2], items[3])
	if err != nil {
		actions.answer.Reply(fmt.Sprint(err))
	} else {
		actions.answer.Reply("Deleted")
	}
}

func help() {

}

// Create creates the ignore handler
func Create(conf *bot.ConfService) bot.Handler {
	myConf := new(configuration)
	conf.Get(myConf)
	log.Notice("Manage only from: %d", myConf.Admins)

	allowed := make(map[int]bool)
	for _, chatID := range myConf.Admins {
		allowed[chatID] = true
	}

	return &handler{allowed}
}
