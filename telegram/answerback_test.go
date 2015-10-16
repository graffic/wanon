package telegram_test

import (
	"testing"

	"github.com/graffic/wanon/mocks"
	"github.com/graffic/wanon/telegram"
)

// TestSend Sends a message back to the chat where it came from.
func TestSend(t *testing.T) {
	message := new(telegram.Message)
	message.Chat.ID = 42

	api := new(mocks.API)
	answeredMessage := new(telegram.Message)
	api.Mock.
		On("SendMessage", &telegram.SendMessage{ChatID: 42, Text: "potato"}).
		Return(answeredMessage, nil)

	answer := telegram.AnswerBack{API: api, Message: message}
	answer.Send("potato")

	api.AssertExpectations(t)
}

// TestReply test it replies to the message
func TestReply(t *testing.T) {
	message := new(telegram.Message)
	message.Chat.ID = 42
	message.MessageID = 13

	api := new(mocks.API)
	answeredMessage := new(telegram.Message)
	api.Mock.
		On("SendMessage", &telegram.SendMessage{
		ChatID:  42,
		Text:    "potato",
		ReplyTo: 13}).
		Return(answeredMessage, nil)

	answer := telegram.AnswerBack{API: api, Message: message}
	answer.Reply("potato")
	api.AssertExpectations(t)
}
