package budget

import (
	"github.com/magleff/gobro/expensefixed"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Budget struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	StartDate     time.Time
	EndDate       time.Time
	ExpensesFixed []expensefixed.ExpenseFixed
	Active        bool
}

func NewBudget(expensesFixed []expensefixed.ExpenseFixed) *Budget {
	instance := new(Budget)
	instance.ExpensesFixed = expensesFixed
	instance.StartDate = time.Now()
	instance.Active = true
	return instance
}
