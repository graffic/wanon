package mocks

import "github.com/graffic/wanon/telegram"
import "github.com/stretchr/testify/mock"

type Request struct {
	mock.Mock
}

func (_m *Request) Call(method string, in interface{}) (*telegram.Response, error) {
	ret := _m.Called(method, in)

	var r0 *telegram.Response
	if rf, ok := ret.Get(0).(func(string, interface{}) *telegram.Response); ok {
		r0 = rf(method, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegram.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(method, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
