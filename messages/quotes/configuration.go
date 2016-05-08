package quotes

import (
	"github.com/graffic/wanon/bot"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("wanon.messages.quotes")

type quoteHandler struct {
	storage quoteStorage
	allowed map[int]bool
}

// IgnoreConfiguration configuration for the ignore message
type configuration struct {
	Allow []int
}

func (handler *quoteHandler) Check(context *bot.MessageContext) int {
	id := context.Message.Chat.ID

	if handler.allowed[id] {
		logger.Debug("%s authorized", context.Message.Text)
		return bot.RouteAccept
	}
	logger.Debug("Not allowed: %d", id)
	return bot.RouteStop
}

func createQuoteHandler(context *bot.BotContext) quoteHandler {
	myConf := new(configuration)
	context.Conf.Get(myConf)
	logger.Notice("Allowing quotes only from: %v", myConf.Allow)

	allowed := make(map[int]bool)
	for _, chatID := range myConf.Allow {
		allowed[chatID] = true
	}

	return quoteHandler{storage: quoteStorage{context.Storage}, allowed: allowed}
}

// Handlers stores handlers
type Handlers interface {
	AddHandler(definition string, givenHandler bot.Handler)
}

// Setup adds the quote handlers to the router
func Setup(handlers Handlers, context *bot.BotContext) {
	handler := createQuoteHandler(context)

	handlers.AddHandler("/addQuote", &addQuote{handler})
	handlers.AddHandler("/rquote", &randomQuote{handler})
	handlers.AddHandler("/status", &quoteStatus{handler})
}
