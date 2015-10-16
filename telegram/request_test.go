package telegram_test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/graffic/wanon/mocks"
	"github.com/graffic/wanon/telegram"
	"github.com/stretchr/testify/assert"
)

func TestCall_Get(t *testing.T) {
	client := new(mocks.HTTPClient)
	request := telegram.NewRequest(client, "http://telegram/")
	body := ioutil.NopCloser(strings.NewReader("{\"ok\": true, \"result\": 4}"))
	httpResponse := &http.Response{Body: body}

	client.Mock.On("Get", "http://telegram/potato").Return(httpResponse, nil)

	out, err := request.Call("potato", nil)

	assert.Nil(t, err)
	assert.True(t, out.Ok)
	client.AssertExpectations(t)
}
