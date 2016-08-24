package database

import (
	"gopkg.in/mgo.v2"
)

type Schema interface {
	Collection(string) *mgo.Collection
}
