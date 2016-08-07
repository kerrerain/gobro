package database_test

import (
	"github.com/magleff/gobro/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

type FakeDatastore struct {
	database.Datastore
}

func TestDatastore(t *testing.T) {
	fakeDatastore := new(FakeDatastore)
	fakeDatastore.BindSession()

	assert.NotNil(t, fakeDatastore.Session, "Should create a session.")
}
