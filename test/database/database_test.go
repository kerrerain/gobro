package database_test

import (
	"github.com/magleff/gobro/database"
	"testing"
)

/*
	Test functions
*/
func TestGetSingletonSession(t *testing.T) {
	databaseMocked := new(DatabaseMocked)
	databaseMocked.On("DialDatabase").Return(FakeSession{})

	database.InitDatabaseWithImpl(databaseMocked)

	database.GetSession()
	database.GetSession()

	databaseMocked.AssertNumberOfCalls(t, "DialDatabase", 1)
}
