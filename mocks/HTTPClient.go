package mocks

import "github.com/stretchr/testify/mock"

import "io"

import "net/http"

type HTTPClient struct {
	mock.Mock
}

func (_m *HTTPClient) Get(url string) (*http.Response, error) {
	ret := _m.Called(url)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string) *http.Response); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *HTTPClient) Post(url string, bodyType string, body io.Reader) (*http.Response, error) {
	ret := _m.Called(url, bodyType, body)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string, string, io.Reader) *http.Response); ok {
		r0 = rf(url, bodyType, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, io.Reader) error); ok {
		r1 = rf(url, bodyType, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
