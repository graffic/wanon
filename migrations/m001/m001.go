package m001

import (
	"fmt"
	"strconv"

	"github.com/graffic/goejdb"
	"github.com/graffic/wanon/messages/manage"
	"github.com/graffic/wanon/migrations"
	"labix.org/v2/mgo/bson"
)

type ejdbMetadata struct {
	Collections []collectionMetadata `bson:"collections"`
}

type collectionMetadata struct {
	Name string `bson:"name"`
}

func migration001(db *goejdb.Ejdb) error {
	metaBytes, err := db.Meta()
	if err != nil {
		return err
	}

	var metadata ejdbMetadata
	bson.Unmarshal(metaBytes, &metadata)

	for _, collInfo := range metadata.Collections {
		collName := collInfo.Name
		_, err2 := strconv.ParseInt(collName, 10, 64)
		if err2 != nil {
			continue
		}

		mover := manage.NewEjdbQuoteMover(db)
		newCollName := "quotes_" + collName

		db.CreateColl(newCollName, nil)
		_, err3 := mover.Move(collName, newCollName)
		if err3 != nil {
			fmt.Println("Error moving", err3)
			return err3
		}
	}

	return nil
}

func init() {
	migrations.Add(1, migration001)
}
