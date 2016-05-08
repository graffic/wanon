package telegram

import (
	"testing"
)

// TestSend Sends a message back to the chat where it came from.
func TestSend(t *testing.T) {
	message := new(Message)
	message.Chat.ID = 42

	api := new(MockAPI)
	answeredMessage := new(Message)
	api.Mock.
		On("SendMessage", &SendMessage{ChatID: 42, Text: "potato"}).
		Return(answeredMessage, nil)

	answer := AnswerBack{API: api, Message: message}
	answer.Send("potato")

	api.AssertExpectations(t)
}

// TestReply test it replies to the message
func TestReply(t *testing.T) {
	message := new(Message)
	message.Chat.ID = 42
	message.MessageID = 13

	api := new(MockAPI)
	answeredMessage := new(Message)
	api.Mock.
		On("SendMessage", &SendMessage{
		ChatID:  42,
		Text:    "potato",
		ReplyTo: 13}).
		Return(answeredMessage, nil)

	answer := AnswerBack{API: api, Message: message}
	answer.Reply("potato")
	api.AssertExpectations(t)
}
