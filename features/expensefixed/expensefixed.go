package expensefixed

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ExpenseFixed struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Date        time.Time
	Description string
	Amount      float32
}
