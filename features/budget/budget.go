package budget

import (
	"github.com/magleff/gobro/features/expense"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Budget struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	StartDate      time.Time
	EndDate        time.Time
	Expenses       []expense.Expense
	InitialBalance float32
	Active         bool
}

func NewBudget(balance float32, initialExpenses []expense.Expense) *Budget {
	instance := new(Budget)
	instance.StartDate = time.Now()
	instance.Active = true
	instance.Expenses = initialExpenses
	instance.InitialBalance = balance
	return instance
}
