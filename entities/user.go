package entities

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID               bson.ObjectId `bson:"_id,omitempty"`
	CurrentAccountId bson.ObjectId `bson:"accountid,omitempty"`
	CurrentBudgetId  bson.ObjectId `bson:"budgetid,omitempty"`
	Name             string
}
