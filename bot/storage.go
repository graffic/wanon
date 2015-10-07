package bot

import (
	"strconv"

	"github.com/mkilling/goejdb"
	"labix.org/v2/mgo/bson"
)

// StorageConfiguration configuration
type StorageConfiguration struct {
	DBName string `yaml:"dbName"`
}

// Storage for quotes
type Storage struct {
	jb *goejdb.Ejdb
}

// NewStorage creates a new storage
func NewStorage(conf *ConfService) (*Storage, error) {
	var storageConf StorageConfiguration
	conf.Get(&storageConf)

	// Create a new database file and open it
	jb, err := goejdb.Open(storageConf.DBName, goejdb.JBOWRITER|goejdb.JBOCREAT)
	if err != nil {
		return nil, err
	}

	storage := new(Storage)
	storage.jb = jb

	return storage, nil
}

// AddQuote to the storage
func (storage *Storage) AddQuote(chat int, quote interface{}) error {
	col, err1 := storage.jb.CreateColl(strconv.Itoa(chat), nil)
	if err1 != nil {
		return err1
	}

	bytes, err2 := bson.Marshal(quote)
	if err2 != nil {
		return err2
	}

	_, err3 := col.SaveBson(bytes)
	if err3 != nil {
		return err3
	}

	return nil
}

// RQuote get a random quote - ToDo
func (storage *Storage) RQuote(chat int) error {
	return nil
}
