package budget

import (
	"github.com/magleff/gobro/features/expense"
	"github.com/shopspring/decimal"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Budget struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	StartDate      time.Time
	EndDate        time.Time
	Expenses       []expense.Expense
	InitialBalance decimal.Decimal
	Active         bool
}

func NewBudget(balance decimal.Decimal, initialExpenses []expense.Expense) *Budget {
	instance := new(Budget)
	instance.StartDate = time.Now()
	instance.Active = true
	instance.Expenses = initialExpenses
	instance.InitialBalance = balance
	return instance
}
