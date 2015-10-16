package mocks

import "github.com/graffic/wanon/telegram"
import "github.com/stretchr/testify/mock"

type API struct {
	mock.Mock
}

func (_m *API) GetMe() (telegram.User, error) {
	ret := _m.Called()

	var r0 telegram.User
	if rf, ok := ret.Get(0).(func() telegram.User); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(telegram.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *API) GetUpdates(_a0 int) []telegram.Update {
	ret := _m.Called(_a0)

	var r0 []telegram.Update
	if rf, ok := ret.Get(0).(func(int) []telegram.Update); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]telegram.Update)
		}
	}

	return r0
}
func (_m *API) ProcessUpdates(_a0 chan *telegram.Message) {
	_m.Called(_a0)
}
func (_m *API) SendMessage(_a0 *telegram.SendMessage) (*telegram.Message, error) {
	ret := _m.Called(_a0)

	var r0 *telegram.Message
	if rf, ok := ret.Get(0).(func(*telegram.SendMessage) *telegram.Message); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegram.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*telegram.SendMessage) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
