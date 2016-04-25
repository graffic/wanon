package bot

import "github.com/graffic/wanon/telegram"

// RouteNothing do nothing in handling
const RouteNothing = 0
const (
	// RouteAccept handle the message
	RouteAccept = 1 << iota
	// RouteStop stop handling more messages after this one
	RouteStop
)

// Message from the telegram router
type Message struct {
	*telegram.Message
	*telegram.AnswerBack
}

// Handler pairs a check and a handle function
type Handler interface {
	Check(messages *Message) int
	Handle(messages *Message)
}

// Router stores handlers for messages
type Router struct {
	handles []Handler
}

// AddHandler adds a handler to the router
func (router *Router) AddHandler(handler Handler) {
	router.handles = append(router.handles, handler)
}

// RouteMessages checks which handler is the destination of a message
func (router *Router) RouteMessages(messages chan *telegram.Message, context *Context) {
	for {
		message := <-messages
		answer := telegram.AnswerBack{API: context.API, Message: message}
		routerMessage := Message{message, &answer}

		for _, handler := range router.handles {

			result := handler.Check(&routerMessage)
			if (result & RouteAccept) > 0 {
				handler.Handle(&routerMessage)
			}
			if (result & RouteStop) > 0 {
				break
			}
		}
	}
}
