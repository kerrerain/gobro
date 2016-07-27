package database

import (
	"gopkg.in/mgo.v2"
)

type Schema interface {
	Collection(string) *mgo.Collection
}

/*
	Implementation
*/

type MgoSchema struct {
	Schema *mgo.Database
}

func (self MgoSchema) Collection(name string) *mgo.Collection {
	return self.Schema.C(name)
}
