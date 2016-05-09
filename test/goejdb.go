package test

import (
	"os"
	"testing"

	"github.com/graffic/goejdb"
)

// VoidFunc a function with no inputs or outputs
type VoidFunc func()

// GoejdbHelper makes easier to test goejdb
type GoejdbHelper struct {
	onCloseFunctions []VoidFunc
	t                *testing.T
	DB               *goejdb.Ejdb
	name             string
}

// NewGoejdbHelper creates a new helper for testing goejdb dbs
func NewGoejdbHelper(t *testing.T, name string) *GoejdbHelper {
	storage, err := goejdb.Open(name, goejdb.JBOWRITER|goejdb.JBOCREAT)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	helper := &GoejdbHelper{t: t, DB: storage, name: name}
	helper.onClose(func() { storage.Del() })
	helper.onClose(func() { os.Remove(name) })

	return helper
}

func (helper *GoejdbHelper) onClose(toExecute VoidFunc) {
	helper.onCloseFunctions = append(helper.onCloseFunctions, toExecute)
}

// Cleanup files
func (helper *GoejdbHelper) Cleanup() {
	for i := len(helper.onCloseFunctions) - 1; i >= 0; i-- {
		helper.onCloseFunctions[i]()
	}
}

// CreateColl creates a coll and deletes it on close
func (helper *GoejdbHelper) CreateColl(collName string) *goejdb.EjColl {
	coll, err := helper.DB.CreateColl(collName, nil)

	if err != nil {
		helper.t.Log(err)
		helper.t.FailNow()
	}

	helper.onClose(func() { helper.DB.RmColl(collName, true) })

	return coll
}
