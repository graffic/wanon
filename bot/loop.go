package bot

import "github.com/graffic/wanon/telegram"

// Router routes messages
type Router interface {
	Route(message *telegram.Message)
}

// MainLoop of the bot to get messages and route them
func MainLoop(channel chan *telegram.Message, router Router) {
	for message := range channel {
		router.Route(message)
	}
}
