package mocks

import "github.com/graffic/wanon/bot"
import "github.com/stretchr/testify/mock"

import "github.com/graffic/wanon/telegram"

type Handler struct {
	mock.Mock
}

func (_m *Handler) Check(_a0 *telegram.Message, _a1 *bot.Context) int {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(*telegram.Message, *bot.Context) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
func (_m *Handler) Handle(_a0 *telegram.Message, _a1 *bot.Context) {
	_m.Called(_a0, _a1)
}
