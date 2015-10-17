package telegram_test

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/graffic/wanon/mocks"
	"github.com/graffic/wanon/telegram"
	"github.com/stretchr/testify/suite"
)

type MockCloser struct {
	io.Reader
	Closed bool
}

func (closer *MockCloser) Close() error {
	closer.Closed = true
	return nil
}

// Normal get request where everything goes ok
type TestCallGetOkSuite struct {
	suite.Suite
	client  mocks.HTTPClient
	body    MockCloser
	request telegram.Request
}

func (suite *TestCallGetOkSuite) SetupTest() {
	suite.client = mocks.HTTPClient{}
	suite.request = telegram.NewRequest(&suite.client, "http://telegram/")

	contents := strings.NewReader("{\"ok\": true, \"result\": 4}")
	suite.body = MockCloser{contents, false}

	httpResponse := &http.Response{Body: &suite.body}
	suite.client.Mock.On("Get", "http://telegram/potato").Return(httpResponse, nil)
}

func (suite *TestCallGetOkSuite) TestRightResponse() {
	out, _ := suite.request.Call("potato", nil)
	suite.True(out.Ok)
}

func (suite *TestCallGetOkSuite) TestNoError() {
	_, err := suite.request.Call("potato", nil)
	suite.Nil(err)
}

func (suite *TestCallGetOkSuite) TestHTTPGetCalled() {
	suite.request.Call("potato", nil)
	suite.client.AssertExpectations(suite.T())
}

func (suite *TestCallGetOkSuite) TestReaderClosed() {
	suite.request.Call("potato", nil)
	suite.True(suite.body.Closed)
}

func TestCall_Get(t *testing.T) {
	suite.Run(t, new(TestCallGetOkSuite))
}

// Error in http.Get
type TestCallGetError struct {
	suite.Suite
	client  mocks.HTTPClient
	request telegram.Request
	err     error
}

func (suite *TestCallGetError) SetupTest() {
	client := mocks.HTTPClient{}
	suite.request = telegram.NewRequest(&suite.client, "http://telegram/")

	suite.err = errors.New("404")
	client.Mock.On("Get", "http://telegram/potato").Return(nil, suite.err)

	suite.client = client
}

func (suite *TestCallGetError) TestNoResponse() {
	out, _ := suite.request.Call("potato", nil)
	suite.Nil(out)
}

func (suite *TestCallGetError) TestError() {
	_, err := suite.request.Call("potato", nil)
	suite.Equal(err, suite.err)
}

func (suite *TestCallGetError) TestHTTPGetCalled() {
	suite.request.Call("potato", nil)
	suite.client.AssertExpectations(suite.T())
}

func TestCall_GetError(t *testing.T) {
	suite.Run(t, new(TestCallGetError))
}

// Error deserializing
type TestCallGetUnmarshalError struct {
	suite.Suite
	client  mocks.HTTPClient
	body    MockCloser
	request telegram.Request
}

func (suite *TestCallGetUnmarshalError) SetupTest() {
	suite.client = mocks.HTTPClient{}
	suite.request = telegram.NewRequest(&suite.client, "http://telegram/")

	contents := strings.NewReader("-- WRONG JSON --")
	suite.body = MockCloser{contents, false}

	httpResponse := &http.Response{Body: &suite.body}
	suite.client.Mock.On("Get", "http://telegram/potato").Return(httpResponse, nil)
}

func (suite *TestCallGetUnmarshalError) TestNoResponse() {
	out, _ := suite.request.Call("potato", nil)
	suite.Nil(out)
}

func (suite *TestCallGetUnmarshalError) TestError() {
	_, err := suite.request.Call("potato", nil)
	suite.Error(err)
}

func (suite *TestCallGetUnmarshalError) TestHTTPGetCalled() {
	suite.request.Call("potato", nil)
	suite.client.AssertExpectations(suite.T())
}

func (suite *TestCallGetUnmarshalError) TestReaderClosed() {
	suite.request.Call("potato", nil)
	suite.True(suite.body.Closed)
}

func TestCall_GetUnmarshalError(t *testing.T) {
	suite.Run(t, new(TestCallGetUnmarshalError))
}
