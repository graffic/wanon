package bot

import (
	"sync"
	"testing"

	"github.com/graffic/wanon/telegram"
)

type testRouter struct {
	msgs []*telegram.Message
}

func (router *testRouter) Route(msg *telegram.Message) {
	router.msgs = append(router.msgs, msg)
}

// Tests the
func TestMainLoop(t *testing.T) {
	channel := make(chan *telegram.Message)
	router := &testRouter{}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		MainLoop(channel, router)
	}()

	var msg telegram.Message
	channel <- &msg
	close(channel)
	wg.Wait()

	if len(router.msgs) != 1 {
		t.Error("Nothing routed")
	}

	if router.msgs[0] != &msg {
		t.Error("Message not routed")
	}
}
