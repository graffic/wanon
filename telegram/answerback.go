package telegram

type messageSender interface {
	SendMessage(message *SendMessage) (*Message, error)
}

// AnswerBack message helper to answer messages
type AnswerBack struct {
	*Message
	sender messageSender
}

// NewAnswerBack creates a new instance
func NewAnswerBack(message *Message, sender messageSender) *AnswerBack {
	return &AnswerBack{message, sender}
}

// Send a message back to the chat where it originated
func (answer *AnswerBack) Send(message string) (*Message, error) {
	return answer.sender.SendMessage(&SendMessage{
		ChatID: answer.Message.Chat.ID,
		Text:   message,
	})
}

// Reply to the sender of the message in the chat it was originated
func (answer *AnswerBack) Reply(message string) (*Message, error) {
	return answer.sender.SendMessage(&SendMessage{
		ChatID:  answer.Message.Chat.ID,
		ReplyTo: answer.Message.MessageID,
		Text:    message,
	})
}
