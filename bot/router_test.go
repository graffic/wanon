package bot

import (
	"testing"

	"github.com/graffic/wanon/telegram"
)

type aloneHandler struct {
	Done int
}

func (handler *aloneHandler) Check(context *MessageContext) int {
	return RouteAccept
}

func (handler *aloneHandler) Handle(context *MessageContext) {
	handler.Done++
}

func (handler *aloneHandler) IsNotCalled() bool {
	return handler.Done == 0
}

type oneHandler struct {
	aloneHandler
	One string
}

func (handler *oneHandler) Handle(context *MessageContext) {
	handler.Done++
	handler.One = context.Params["paramOne"]
}

type twoHandler struct {
	oneHandler
	Two string
}

func (handler *twoHandler) Handle(context *MessageContext) {
	handler.Done++
	handler.One = context.Params["paramOne"]
	handler.Two = context.Params["paramTwo"]
}

type testHandler interface {
	IsNotCalled() bool
}

func allNotCalled(t *testing.T, handlers []testHandler) {
	for _, handler := range handlers {
		if !handler.IsNotCalled() {
			t.Error("Handler executed", handler)
		}
	}
}

func setupRouter() (*Routes, *aloneHandler, *oneHandler, *twoHandler) {
	routes := new(Routes)

	alone := new(aloneHandler)
	routes.AddHandler("/potato", alone)

	one := new(oneHandler)
	routes.AddHandler("/potato :paramOne", one)

	two := new(twoHandler)
	routes.AddHandler("/potato :paramOne :paramTwo", two)

	return routes, alone, one, two
}

func checkExecutions(t *testing.T, done int) {
	if done == 0 {
		t.Error("Not executed")
	}

	if done > 1 {
		t.Error("Executed more than once")
	}
}

func TestAlone(t *testing.T) {
	router, alone, other1, other2 := setupRouter()
	message := telegram.Message{Text: "/potato"}
	router.Route(&message)

	checkExecutions(t, alone.Done)
	allNotCalled(t, []testHandler{other1, other2})
}

func TestOne(t *testing.T) {
	router, other1, one, other2 := setupRouter()
	message := telegram.Message{Text: "/potato one"}
	router.Route(&message)

	checkExecutions(t, one.Done)
	allNotCalled(t, []testHandler{other1, other2})

	if one.One != "one" {
		t.Error("One parameter is not correct", one.One)
	}
}

func TestTwo(t *testing.T) {
	router, other1, other2, two := setupRouter()
	message := telegram.Message{Text: "/potato one two"}
	router.Route(&message)

	checkExecutions(t, two.Done)
	allNotCalled(t, []testHandler{other1, other2})

	if two.One != "one" {
		t.Error("One parameter is not correct", two.One)
	}
	if two.Two != "two" {
		t.Error("Two parameter is not correct", two.Two)
	}
}
