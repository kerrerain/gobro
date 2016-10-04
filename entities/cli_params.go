package entities

import (
	"gopkg.in/mgo.v2/bson"
)

type CliParams struct {
	CurrentAccountId bson.ObjectId `bson:"currentaccountid,omitempty"`
	CurrentBudgetId  bson.ObjectId `bson:"currentbudgetid,omitempty"`
}
