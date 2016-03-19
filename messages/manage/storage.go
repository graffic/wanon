package manage

import (
	"fmt"

	"github.com/graffic/wanon/bot"
	"github.com/graffic/wanon/messages/quotes"

	"labix.org/v2/mgo/bson"
)

type managerStorage struct {
	db bot.Storage
}

type ejdbMetadata struct {
	File        string               `bson:"file"`
	Collections []collectionMetadata `bson:"collections"`
}

type collectionMetadata struct {
	Name    string `bson:"name"`
	Records int    `bson:"records"`
}

func (storage *managerStorage) Chats() (*[]collectionMetadata, error) {
	bsonData, err := storage.db.Meta()
	if err != nil {
		return nil, err
	}

	var metadata ejdbMetadata
	bson.Unmarshal(bsonData, &metadata)

	return &metadata.Collections, nil
}

func (storage *managerStorage) Delete(chat string, oid string) error {
	coll, err := storage.db.GetColl(chat)
	if err != nil {
		return err
	}
	coll.RmBson(oid)
	return nil
}

func (storage *managerStorage) List(chat string, amountToSkip int) (*[]quotes.Quote, error) {
	col, err := storage.db.GetColl(chat)
	if err != nil {
		return nil, err
	}
	query, err := storage.db.CreateQuery("{}")
	defer query.Del()
	if err != nil {
		return nil, err
	}
	hint := fmt.Sprintf(`{"$orderby":{"when":1}, "$skip": %d}`, amountToSkip)
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
