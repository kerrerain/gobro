package database

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

/*
	Test objects
*/

type DatabaseMocked struct {
	mock.Mock
}

func (m *DatabaseMocked) DialDatabase() Session {
	args := m.Called()
	return args.Get(0).(Session)
}

/*
	Test functions
*/
func TestGetSingletonSession(t *testing.T) {
	InitDatabase()

	databaseMocked := new(DatabaseMocked)
	databaseMocked.On("DialDatabase").Return(FakeSession{})
	database = databaseMocked

	GetSession()
	GetSession()

	databaseMocked.AssertNumberOfCalls(t, "DialDatabase", 1)
}
