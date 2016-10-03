package entities

import (
	"github.com/shopspring/decimal"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Budget struct {
	ID               bson.ObjectId `bson:"_id,omitempty"`
	AccountId        bson.ObjectId `bson:"accountid,omitempty"`
	UserId           bson.ObjectId `bson:"userid,omitempty"`
	StartDate        time.Time
	LastModification time.Time
	Expenses         []Expense
	InitialBalance   decimal.Decimal
	Active           bool
}
