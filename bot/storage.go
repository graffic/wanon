package bot

import "github.com/mkilling/goejdb"

// StorageConfiguration configuration
type StorageConfiguration struct {
	DBName string `yaml:"dbName"`
}

// Storage for quotes
type Storage interface {
	CreateColl(colname string, opts *goejdb.EjCollOpts) (*goejdb.EjColl, *goejdb.EjdbError)
	CreateQuery(query string, queries ...string) (*goejdb.EjQuery, *goejdb.EjdbError)
	GetColl(colname string) (*goejdb.EjColl, *goejdb.EjdbError)
	Meta() ([]byte, *goejdb.EjdbError)
}

// NewStorage creates a new storage
func NewStorage(conf *ConfService) (Storage, error) {
	var storageConf StorageConfiguration
	conf.Get(&storageConf)

	// Create a new database file and open it
	db, err := goejdb.Open(storageConf.DBName, goejdb.JBOWRITER|goejdb.JBOCREAT)
	if err != nil {
		return nil, err
	}

	return db, nil
}
