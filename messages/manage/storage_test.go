package manage

import (
	"os"
	"testing"

	"labix.org/v2/mgo/bson"

	"github.com/graffic/wanon/messages/quotes"
	"github.com/graffic/wanon/mocks"
	"github.com/mkilling/goejdb"
	"github.com/stretchr/testify/assert"
)

// TestChats Return only collection with name as integer
func TestChats(t *testing.T) {
	var storage mocks.Storage
	manager := managerStorage{db: &storage}

	metadata := ejdbMetadata{
		"file",
		[]collectionMetadata{
			collectionMetadata{"test1", 1},
			collectionMetadata{"12345", 2},
		},
	}
	bsonBytes, _ := bson.Marshal(metadata)
	storage.Mock.On("Meta").Return(bsonBytes, nil)

	names, _ := manager.Chats()

	assert.Equal(t, &metadata.Collections, names)
}

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

	manager := managerStorage{db: storage}

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

	manager := managerStorage{db: storage}

	res, err2 := manager.List("12345", 10)
	if err2 != nil {
		t.Error(err2)
	}

	assert.Equal(t, 10, len(*res))
}
