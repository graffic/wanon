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
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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

// Chat payload
type Chat struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Message payload
type Message struct {
	MessageID      int      `json:"message_id"`
	From           User     `json:"from"`
	Text           string   `json:"text"`
	Date           int      `json:"date"`
	Chat           Chat     `json:"chat"`
	ReplyToMessage *Message `json:"reply_to_message"`
}

// SendMessage action parameter
type SendMessage struct {
	ChatID    int    `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
	ReplyTo   int    `json:"reply_to_message_id"`
}

// ReplyKeyboardMarkup custom keyboard with reply options
type ReplyKeyboardMarkup struct {
	Keyboard        [][]string `json:"keyboard"`
	ResizeKeyboard  bool       `json:"resize_keyboard"`
	OneTimeKeyboard bool       `json:"one_time_keyboard"`
	Selective       bool       `json:"selective"`
}

// ReplyKeyboardHide hide the current keyboard
type ReplyKeyboardHide struct {
	HideKeyboard bool `json:"hide_keyboard"`
	Selective    bool `json:"selective"`
}

// ForceReply show reply interface
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

// GetUpdates action parameter
type GetUpdates struct {
	Offset  int `json:"offset"`
	Timeout int `json:"timeout"`
}
