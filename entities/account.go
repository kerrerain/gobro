package entities

import (
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	UserId bson.ObjectId `bson:"userid,omitempty"`
	Name   string
	Label  string
	Active bool
}
