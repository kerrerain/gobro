package entities

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	CliParams CliParams
}
