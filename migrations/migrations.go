package migrations

import (
	"sort"

	"github.com/graffic/goejdb"
	"github.com/op/go-logging"
)

// Migration function
type Migration func(db *goejdb.Ejdb) error

// Migrations type to hold all migrations
type Migrations []*metadata

var migrations Migrations

type metadata struct {
	order     int
	migration Migration
}

func (m Migrations) Len() int {
	return len(m)
}

func (m Migrations) Less(i, j int) bool {
	return m[i].order < m[j].order
}

func (m Migrations) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func add(order int, migration Migration) {
	migrations = append(migrations, &metadata{order, migration})
}

// Run migrations in order
func Run(db *goejdb.Ejdb) error {
	logger := logging.MustGetLogger("wanon.migrations")
	sort.Sort(migrations)

	versions, err := NewSchemaVersion(db)
	if err != nil {
		return err
	}

	latest, err := versions.GetLatest()
	if err != nil {
		return err
	}

	for _, item := range migrations {
		if item.order <= latest {
			continue
		}

		err = item.migration(db)
		if err != nil {
			logger.Critical("Error in migration", item.order, err)
			return err
		}

		err2 := versions.Add(item.order)
		if err2 != nil {
			return err
		}
	}
	return nil
}
