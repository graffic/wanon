package bot

import (
	"strings"

	"github.com/graffic/wanon/telegram"
)

// RouteNothing do nothing in handling
const RouteNothing = 0
const (
	// RouteAccept handle the message
	RouteAccept = 1 << iota
	// RouteStop stop handling more messages after this one
	RouteStop
)

// MessageContext for each received message
type MessageContext struct {
	Params  map[string]string
	Message *telegram.AnswerBack
}

// Handler pairs a check and a handle function
type Handler interface {
	Check(messages *MessageContext) int
	Handle(messages *MessageContext)
}

// Routes stores handlers for messages
type Routes struct {
	Bot      BotContext
	handlers []handler
}

type handler struct {
	elements []string
	actions  Handler
}

// AddHandler for incoming messages
func (router *Routes) AddHandler(definition string, givenHandler Handler) {
	elements := strings.Split(definition, " ")
	router.handlers = append(router.handlers, handler{elements, givenHandler})
}

// Route from a telegram message channel to the respective handler
func (router *Routes) Route(message *telegram.Message) {
	var elements []string
	for _, item := range strings.Split(message.Text, " ") {
		if item == "" {
			continue
		}
		elements = append(elements, item)
	}

	count := len(elements)
	answer := telegram.AnswerBack{API: router.Bot.API, Message: message}
	messageContext := MessageContext{map[string]string{}, &answer}

	for _, handler := range router.handlers {
		if count != len(handler.elements) {
			continue
		}
		matched := true
		for index, item := range handler.elements {
			givenItem := elements[index]
			switch {
			case strings.Index(item, ":") == 0:
				messageContext.Params[item[1:]] = givenItem
			case item != givenItem:
				matched = false
			}
		}
		if !matched {
			continue
		}
		check := handler.actions.Check(&messageContext)

		switch {
		case check == RouteAccept:
			handler.actions.Handle(&messageContext)
			return
		case check == RouteStop:
			return
		}
	}
}

// Router routes messages
type Router interface {
	Route(message *telegram.Message)
}

// MainLoop of the bot to get messages and route them
func MainLoop(channel chan *telegram.Message, router Router) {
	for {
		message := <-channel
		router.Route(message)
	}
}
