package quotes

import (
	"fmt"
	"math/rand"

	"github.com/graffic/goejdb"

	"labix.org/v2/mgo/bson"
)

type quoteStorage struct {
	ejdb *goejdb.Ejdb
}

// Quote what a quote contains
type Quote struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	SaidBy  string
	AddedBy string
	When    int
	What    string
}

func (storage *quoteStorage) getQuoteColl(chat int) (*goejdb.EjColl, error) {
	collName := fmt.Sprintf("quotes_%d", chat)

	col, err := storage.ejdb.CreateColl(collName, nil)
	if err != nil {
		return nil, err
	}
	return col, nil
}

// AddQuote to the storage
func (storage *quoteStorage) AddQuote(chat int, quote *Quote) error {
	col, err1 := storage.getQuoteColl(chat)
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
	col, err := storage.getQuoteColl(chat)
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
	logger.Debug("Entries: %d, Skip: %d", count, amountToSkip)

	skip := fmt.Sprintf(`{"$skip": %d}`, amountToSkip)
	query, err := storage.ejdb.CreateQuery("{}")
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
