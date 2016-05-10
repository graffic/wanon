package migrations

import (
	"testing"

	"github.com/graffic/wanon/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSchemaVersion struct {
	suite.Suite
	helper        *test.GoejdbHelper
	schemaVersion *SchemaVersion
}

func (suite *TestSchemaVersion) SetupTest() {
	suite.helper = test.NewGoejdbHelper(suite.T(), "migrations_internal_test")
	suite.schemaVersion, _ = NewSchemaVersion(suite.helper.DB)
}

func (suite *TestSchemaVersion) TearDownTest() {
	suite.helper.DB.RmColl(versionsCollectionName, true)
	suite.helper.Cleanup()
}

func (suite *TestSchemaVersion) TestLastVersionWhenEmpty() {
	latest, _ := suite.schemaVersion.GetLatest()
	assert.Equal(suite.T(), 0, latest)
}

func (suite *TestSchemaVersion) TestLastVersionAfterAdd() {
	version := suite.schemaVersion
	version.Add(5)
	version.Add(10)
	version.Add(2)
	latest, _ := version.GetLatest()
	assert.Equal(suite.T(), 10, latest)
}

func TestSuite_TestSchemaVersion(t *testing.T) {
	suite.Run(t, new(TestSchemaVersion))
}
