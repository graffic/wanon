package telegram

import "github.com/stretchr/testify/mock"

// This is an autogenerated mock type for the API type
type MockAPI struct {
	mock.Mock
}

// GetMe provides a mock function with given fields:
func (_m *MockAPI) GetMe() (User, error) {
	ret := _m.Called()

	var r0 User
	if rf, ok := ret.Get(0).(func() User); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpdates provides a mock function with given fields: _a0
func (_m *MockAPI) GetUpdates(_a0 int) []Update {
	ret := _m.Called(_a0)

	var r0 []Update
	if rf, ok := ret.Get(0).(func(int) []Update); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Update)
		}
	}

	return r0
}

// ProcessUpdates provides a mock function with given fields: _a0
func (_m *MockAPI) ProcessUpdates(_a0 chan *Message) {
	_m.Called(_a0)
}

// SendMessage provides a mock function with given fields: _a0
func (_m *MockAPI) SendMessage(_a0 *SendMessage) (*Message, error) {
	ret := _m.Called(_a0)

	var r0 *Message
	if rf, ok := ret.Get(0).(func(*SendMessage) *Message); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*SendMessage) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
