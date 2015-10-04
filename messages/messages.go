package messages

import (
	"github.com/graffic/wanon/messages/addquote"
	"github.com/graffic/wanon/messages/ignorechat"
)

// CreateIgnore create ignore handler
var CreateIgnore = ignorechat.Create

// CreateAddQuote /addquote handler
var CreateAddQuote = addquote.Create
