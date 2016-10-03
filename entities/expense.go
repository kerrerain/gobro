package entities

import (
	"github.com/shopspring/decimal"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Expense struct {
	ID          bson.ObjectId   `json:"_id,omitempty" bson:"_id,omitempty"`
	Date        time.Time       `json:"date" bson:"date"`
	Description string          `json:"description" bson:"description"`
	Amount      decimal.Decimal `json:"amount" bson:"amount"`
	Checked     bool            `json:"checked" bson:"checked"`
}
