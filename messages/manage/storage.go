package manage

import (
	"fmt"

	"github.com/graffic/goejdb"
	"github.com/graffic/wanon/messages/quotes"
	"labix.org/v2/mgo/bson"
)

type manageStorage struct {
	ejdb *goejdb.Ejdb
}

type ejdbMetadata struct {
	File        string               `bson:"file"`
	Collections []collectionMetadata `bson:"collections"`
}

type collectionMetadata struct {
	Name    string `bson:"name"`
	Records int    `bson:"records"`
}

func (storage *manageStorage) Move(from string, to string) (int, error) {
	fromCol, err := storage.ejdb.GetColl(from)
	if err != nil {
		return 0, err
	}

	toCol, err := storage.ejdb.GetColl(to)
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

	storage.ejdb.RmColl(from, true)

	return amount, nil
}

func (storage *manageStorage) Chats() (*[]collectionMetadata, error) {
	bsonData, err := storage.ejdb.Meta()
	if err != nil {
		return nil, err
	}

	var metadata ejdbMetadata
	bson.Unmarshal(bsonData, &metadata)

	var results []collectionMetadata
	for _, collInfo := range metadata.Collections {
		results = append(results, collInfo)
	}

	return &metadata.Collections, nil
}

// Delete a quote form the specific channel
func (storage *manageStorage) Delete(chat string, oid string) error {
	coll, err := storage.ejdb.GetColl(chat)
	if err != nil {
		return err
	}
	coll.RmBson(oid)
	return nil
}

func (storage *manageStorage) List(chat string, amountToSkip int) (*[]quotes.Quote, error) {
	col, err := storage.ejdb.GetColl(chat)
	if err != nil {
		return nil, err
	}
	query, err := storage.ejdb.CreateQuery("{}")
	defer query.Del()
	if err != nil {
		return nil, err
	}
	hint := fmt.Sprintf(`{"$orderby":{"when":1}, "$skip": %d, "$max": 10}`, amountToSkip)
	query.SetHints(hint)

	results, err := query.Execute(col)
	if err != nil {
		return nil, err
	}

	res := []quotes.Quote{}
	for _, quoteBytes := range results {
		var quote quotes.Quote
		bson.Unmarshal(quoteBytes, &quote)
		res = append(res, quote)
	}

	return &res, nil
}
