package m001

import (
	"testing"

	"github.com/graffic/wanon/migrations"
	"github.com/graffic/wanon/test"
	"github.com/stretchr/testify/assert"
)

func Test001Present(t *testing.T) {
	assert.True(t, migrations.IsPresent(1))
}

func Test001MovesData(t *testing.T) {
	helper := test.NewGoejdbHelper(t, "migration_001")
	defer helper.Cleanup()

	coll := helper.CreateColl("-12345")
	coll.SaveJson("{}")

	err := migration001(helper.DB)
	if err != nil {
		t.Error(err)
	}

	coll = helper.CreateColl("quotes_-12345")
	amount, _ := coll.Count("{}")

	if amount != 1 {
		t.Error("Migration didn't move the data")
	}
}

func Test001RemovesOld(t *testing.T) {
	helper := test.NewGoejdbHelper(t, "migration_001")
	defer helper.Cleanup()

	coll := helper.CreateColl("12345")
	coll.SaveJson("{}")

	migration001(helper.DB)

	coll, err := helper.DB.GetColl("12345")
	// To delte the file after the test
	defer helper.CreateColl("quotes_12345")

	if err == nil {
		t.Error("Old collection still exists")
	}
}

func Test001DoesntMoveNonQuoteColls(t *testing.T) {
	helper := test.NewGoejdbHelper(t, "migration_001")
	defer helper.Cleanup()

	coll := helper.CreateColl("potato_123")
	coll.SaveJson("{}")

	migration001(helper.DB)

	coll = helper.CreateColl("potato_123")
	amount, _ := coll.Count("{}")

	if amount != 1 {
		t.Error("It moved something it shouldn't")
	}
}
