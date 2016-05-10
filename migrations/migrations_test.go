package migrations

import (
	"errors"
	"testing"

	"github.com/graffic/goejdb"
	"github.com/graffic/wanon/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestMigrations struct {
	suite.Suite
	migrations *Migrations
	result     string
	dbHelper   *test.GoejdbHelper
	given      [2]*goejdb.Ejdb
}

func (suite *TestMigrations) SetupTest() {
	suite.dbHelper = test.NewGoejdbHelper(suite.T(), "migrations_test")

	migrations = Migrations{}
	add(10, func(db *goejdb.Ejdb) error {
		suite.result = suite.result + " world"
		suite.given[0] = db
		return nil
	})
	add(4, func(db *goejdb.Ejdb) error {
		suite.result = "hello"
		suite.given[1] = db
		return nil
	})
}

func (suite *TestMigrations) TearDownTest() {
	suite.dbHelper.DB.RmColl(versionsCollectionName, true)
	suite.dbHelper.Cleanup()
}

func (suite *TestMigrations) TestOrderAndExecution() {
	Run(suite.dbHelper.DB)

	assert.Equal(suite.T(), "hello world", suite.result)
}

func (suite *TestMigrations) TestGivenDatabase() {
	db := suite.dbHelper.DB
	Run(db)

	assert.Equal(suite.T(), [2]*goejdb.Ejdb{db, db}, suite.given)
}

func (suite *TestMigrations) TestFailOnError() {
	migrationError := errors.New("Fake error")
	add(5, func(db *goejdb.Ejdb) error {
		return migrationError
	})

	err := Run(suite.dbHelper.DB)
	assert.Equal(suite.T(), migrationError, err)
	assert.Equal(suite.T(), "hello", suite.result)
}

func (suite *TestMigrations) TestShouldNotRerunPrevious() {
	Run(suite.dbHelper.DB)
	suite.given[0], suite.given[1] = nil, nil
	Run(suite.dbHelper.DB)

	assert.Equal(suite.T(), [2]*goejdb.Ejdb{nil, nil}, suite.given)
}

func TestSuite_TestMigrations(t *testing.T) {
	suite.Run(t, new(TestMigrations))
}
