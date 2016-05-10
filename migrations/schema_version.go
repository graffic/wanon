package migrations

import (
	"time"

	"github.com/graffic/goejdb"
	"labix.org/v2/mgo/bson"
)

const versionsCollectionName = "schema_version"

// SchemaVersion of an ejdb database
type SchemaVersion struct {
	collection *goejdb.EjColl
	latest     func() (*[]byte, *goejdb.EjdbError)
}

type versionItem struct {
	Version int   `bson:"version"`
	When    int64 `bson:"when"`
}

// NewSchemaVersion creates a schema version service
func NewSchemaVersion(db *goejdb.Ejdb) (*SchemaVersion, error) {
	col, err := db.GetColl(versionsCollectionName)
	if err != nil {
		col, err = db.CreateColl(versionsCollectionName, nil)
		if err != err {
			return nil, err
		}

		col.SetIndex("version", goejdb.JBIDXNUM)
	}

	query, err := db.CreateQuery("{}")
	if err != nil {
		return nil, err
	}
	query.SetHints(`{"$orderby": {"version": -1}}`)

	return &SchemaVersion{col, func() (*[]byte, *goejdb.EjdbError) {
		return query.ExecuteOne(col)
	}}, nil
}

// Add a version
func (sv *SchemaVersion) Add(version int) error {
	item := versionItem{version, time.Now().Unix()}
	bytes, err1 := bson.Marshal(item)
	if err1 != nil {
		return err1
	}

	_, err2 := sv.collection.SaveBson(bytes)
	if err2 != nil {
		return err2
	}

	return nil
}

// GetLatest get the latest version
func (sv *SchemaVersion) GetLatest() (int, error) {
	data, err := sv.latest()
	if err != nil {
		return 0, err
	}

	if data == nil {
		return 0, nil
	}

	var item versionItem
	err2 := bson.Unmarshal(*data, &item)
	if err2 != nil {
		return 0, err2
	}

	return item.Version, nil
}
