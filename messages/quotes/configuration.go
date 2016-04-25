package quotes

import (
	"strings"

	"github.com/graffic/wanon/bot"
)

type quoteHandler struct {
	storage quoteStorage
	allowed map[int]bool
}

// IgnoreConfiguration configuration for the ignore message
type configuration struct {
	Allow []int
}

func (handler *quoteHandler) check(command string, message *bot.Message) int {
	isCommand := strings.Index(message.Text, command) == 0

	if handler.allowed[message.Chat.ID] && isCommand {
		log.Debug("%s requested", command)
		return bot.RouteAccept
	}
	if isCommand {
		log.Debug("Not allowed: %d", message.Chat.ID)
		return bot.RouteStop
	}
	return bot.RouteNothing
}

func createQuoteHandler(context *bot.Context) quoteHandler {
	myConf := new(configuration)
	context.Conf.Get(myConf)
	log.Notice("Allowing quotes only from: %v", myConf.Allow)

	allowed := make(map[int]bool)
	for _, chatID := range myConf.Allow {
		allowed[chatID] = true
	}

	return quoteHandler{storage: quoteStorage{context.Storage}, allowed: allowed}
}

// Setup adds the quote handlers to the router
func Setup(router *bot.Router, context *bot.Context) {
	handler := createQuoteHandler(context)

	router.AddHandler(&addQuote{handler})
	router.AddHandler(&randomQuote{handler})
	router.AddHandler(&quoteStatus{handler})
}
