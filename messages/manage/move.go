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

// NewEjdbQuoteMover creates a new quote mover
func NewEjdbQuoteMover(db *goejdb.Ejdb) *EjdbQuoteMover {
	return &EjdbQuoteMover{db}
}

// EjdbQuoteMover moves items from a collection to another
type EjdbQuoteMover struct {
	db *goejdb.Ejdb
}

// Move from coll to coll
func (mover *EjdbQuoteMover) Move(from string, to string) (int, error) {
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
		_, err := toCol.SaveBson(quote)
		if err != nil {
			return 0, err
		}
	}

	mover.db.RmColl(from, true)

	return amount, nil
}
