package manage

import (
	"os"
	"testing"

	"labix.org/v2/mgo/bson"

	"github.com/graffic/goejdb"
	"github.com/graffic/wanon/messages/quotes"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationList_Sorting(t *testing.T) {
	storage, err := goejdb.Open("TestIntegrationList", goejdb.JBOWRITER|goejdb.JBOCREAT)
	if err != nil {
		t.Error(err)
	}
	defer storage.Del()
	defer os.Remove("TestIntegrationList")

	chatColl, err := storage.CreateColl("12345", nil)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove("TestIntegrationList_12345")

	for i := 11; i > 0; i-- {
		bsonBytes, _ := bson.Marshal(quotes.Quote{When: i})
		chatColl.SaveBson(bsonBytes)
	}

	manager := manageStorage{storage}

	res, err2 := manager.List("12345", 10)
	if err2 != nil {
		t.Error(err2)
	}

	assert.Equal(t, 11, (*res)[0].When)
}

func TestIntegrationList_Pagination(t *testing.T) {
	storage, err := goejdb.Open("TestIntegrationList", goejdb.JBOWRITER|goejdb.JBOCREAT)
	if err != nil {
		t.Error(err)
	}
	defer storage.Del()
	defer os.Remove("TestIntegrationList")

	chatColl, err := storage.CreateColl("12345", nil)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove("TestIntegrationList_12345")

	for i := 25; i > 0; i-- {
		bsonBytes, _ := bson.Marshal(quotes.Quote{When: i})
		chatColl.SaveBson(bsonBytes)
	}

	manager := manageStorage{storage}

	res, err2 := manager.List("12345", 10)
	if err2 != nil {
		t.Error(err2)
	}

	assert.Equal(t, 10, len(*res))
}
