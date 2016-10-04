package database

import (
	"gopkg.in/mgo.v2"
)

type MgoSchema struct {
	Schema *mgo.Database
}

func (self MgoSchema) Collection(name string) *mgo.Collection {
	return self.Schema.C(name)
}

func (self MgoSchema) Run(cmd interface{}, result interface{}) error {
	return self.Schema.Run(cmd, result)
}

func (self MgoSchema) DropDatabase() error {
	return self.Schema.DropDatabase()
}
