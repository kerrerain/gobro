package database

import (
	"gopkg.in/mgo.v2"
)

type Schema interface {
	Collection(string) *mgo.Collection
	Run(cmd interface{}, result interface{}) error
	DropDatabase() error
}
