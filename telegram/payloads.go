package telegram

import "encoding/json"

// Response payload
type Response struct {
	Ok     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`
}

// User payload
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// GroupChat payload
type GroupChat struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Update payload
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

// UserAndGroup embeds user and group together
type UserAndGroup struct {
	User
	GroupChat
	ID int `json:"id"`
}

// Message payload
type Message struct {
	MessageID      int          `json:"message_id"`
	From           User         `json:"from"`
	Text           string       `json:"text"`
	Date           int          `json:"date"`
	Chat           UserAndGroup `json:"chat"`
	ReplyToMessage *Message     `json:"reply_to_message"`
}
