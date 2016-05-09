// Package manage move quotes from one chat to another
package manage

import (
	"fmt"

	"github.com/graffic/goejdb"
	"github.com/graffic/wanon/bot"
)

type quoteMover interface {
	Move(from string, to string) (int, error)
}

// MoveHandler for the move command
type moveHandler struct {
	*manageHandler
	storage quoteMover
}

// Handle the move action
func (handler *moveHandler) Handle(context *bot.MessageContext) {
	from := context.Params["from"]
	to := context.Params["to"]
	amount, err := handler.storage.Move(from, to)

	if err != nil {
		logger.Errorf("%v", err)
		return
	}

	context.Message.Reply(fmt.Sprintf("Moved %d quotes", amount))
}

type ejdbQuoteMover struct {
	db *goejdb.Ejdb
}

func (mover *ejdbQuoteMover) Move(from string, to string) (int, error) {
	fromCol, err := mover.db.GetColl(from)
	if err != nil {
		return 0, err
	}

	toCol, err := mover.db.GetColl(to)
	if err != nil {
		return 0, err
	}

	amount, err := fromCol.Count("{}")
	if err != nil {
		return 0, err
	}

	quotes, err := fromCol.Find("{}")
	if err != nil {
		return 0, err
	}

	for _, quote := range quotes {
		toCol.SaveBson(quote)
	}

	mover.db.RmColl(from, true)

	return amount, nil
}
