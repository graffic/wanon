package bot

import "github.com/graffic/goejdb"

// StorageConfiguration configuration
type StorageConfiguration struct {
	DBName string `yaml:"dbName"`
}

// NewStorage creates a new storage
func NewStorage(conf *ConfService) (*goejdb.Ejdb, error) {
	var storageConf StorageConfiguration
	conf.Get(&storageConf)

	// Create a new database file and open it
	db, err := goejdb.Open(storageConf.DBName, goejdb.JBOWRITER|goejdb.JBOCREAT)
	if err != nil {
		return nil, err
	}

	return db, nil
}
