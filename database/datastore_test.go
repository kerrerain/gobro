package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type FakeDatastore struct {
	Datastore
}

func TestDatastore(t *testing.T) {
	fakeDatastore := new(FakeDatastore)
	fakeDatastore.BindSession()

	assert.NotNil(t, fakeDatastore.Session, "Should create a session.")
}
