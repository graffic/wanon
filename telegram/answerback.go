package telegram

// AnswerBack message helper to answer messages
type AnswerBack struct {
	*API
	*Message
}

// Send a message back to the chat where it originated
func (answer *AnswerBack) Send(message string) (*Message, error) {
	return answer.API.SendMessage(&SendMessage{
		ChatID: answer.Message.Chat.ID,
		Text:   message,
	})
}

// Reply to the sender of the message in the chat it was originated
func (answer *AnswerBack) Reply(message string) (*Message, error) {
	return answer.API.SendMessage(&SendMessage{
		ChatID:  answer.Message.Chat.ID,
		ReplyTo: answer.Message.MessageID,
		Text:    message,
	})
}
