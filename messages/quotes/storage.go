package quotes

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/graffic/wanon/bot"
	"labix.org/v2/mgo/bson"
)

type quoteStorage struct {
	*bot.Storage
}

// Quote what a quote contains
type Quote struct {
	SaidBy  string
	AddedBy string
	When    string
	What    string
}

// AddQuote to the storage
func (storage *quoteStorage) AddQuote(chat int, quote *Quote) error {
	col, err1 := storage.DB.CreateColl(strconv.Itoa(chat), nil)
	if err1 != nil {
		return err1
	}

	bytes, err2 := bson.Marshal(quote)
	if err2 != nil {
		return err2
	}

	_, err1 = col.SaveBson(bytes)
	if err1 != nil {
		return err1
	}

	return nil
}

// RQuote get a random quote - ToDo
func (storage *quoteStorage) RQuote(chat int) (*Quote, error) {
	col, err := storage.DB.CreateColl(strconv.Itoa(chat), nil)
	if err != nil {
		return nil, err
	}

	count, err := col.Count("{}")
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, nil
	}

	amountToSkip := rand.Intn(count)

	l.Debug("Entries: %d, Skip: %d", count, amountToSkip)

	skip := fmt.Sprintf(`{"$skip": %d}`, amountToSkip)
	query, err := storage.DB.CreateQuery("{}")
	defer query.Del()
	if err != nil {
		return nil, err
	}
	query.SetHints(skip)

	result, err := query.ExecuteOne(col)
	if err != nil {
		return nil, err
	}

	var quote Quote
	bson.Unmarshal(*result, &quote)

	return &quote, nil
}
