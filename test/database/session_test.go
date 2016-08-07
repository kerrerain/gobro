package database_test

import (
	"github.com/magleff/gobro/database"
)

type FakeSession struct{}

func (self FakeSession) Clone() database.Session {
	return database.MgoSession{}
}

func (self FakeSession) Close() {}

func (self FakeSession) Schema(name string) database.Schema {
	return database.MgoSchema{}
}

func (self FakeSession) DefaultSchema() database.Schema {
	return database.MgoSchema{}
}
