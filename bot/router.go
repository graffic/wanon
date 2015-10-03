package bot

import (
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

// HandlerCheck Checks if a message is for the handler
type HandlerCheck func(*telegram.Message) int

// Handle executes the action
type Handle func(*telegram.Message)

// Handler pairs a check and a handle function
type Handler struct {
	check  HandlerCheck
	handle Handle
}

// NewHandler creates a Handler struct
func NewHandler(check HandlerCheck, handle Handle) Handler {
	return Handler{check, handle}
}

// Router stores handlers for messages
type Router struct {
	handles []*Handler
}

// AddHandler adds a handler to the router
func (router *Router) AddHandler(handler *Handler) {
	router.handles = append(router.handles, handler)
}

// RouteMessages checks which handler is the destination of a message
func (router *Router) RouteMessages(messages chan *telegram.Message) {
	for {
		message := <-messages

		for _, handler := range router.handles {
			result := handler.check(message)
			if (result & RouteAccept) > 0 {
				handler.handle(message)
			}
			if (result & RouteStop) > 0 {
				break
			}
		}
	}
}
