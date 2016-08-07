package database_test

import (
	"github.com/magleff/gobro/database"
	"github.com/stretchr/testify/mock"
)

type DatabaseMocked struct {
	mock.Mock
}

func (m *DatabaseMocked) DialDatabase() database.Session {
	args := m.Called()
	return args.Get(0).(database.Session)
}
