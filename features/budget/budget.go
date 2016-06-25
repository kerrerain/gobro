package budget

import (
	"github.com/magleff/gobro/features/expense"
	"github.com/magleff/gobro/features/expensefixed"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Budget struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	StartDate time.Time
	EndDate   time.Time
	Expenses  []expense.Expense
	Active    bool
}

func NewBudget(expensesFixed []expensefixed.ExpenseFixed) *Budget {
	instance := new(Budget)
	instance.Expenses = convertExpensesFixed(expensesFixed)
	instance.StartDate = time.Now()
	instance.Active = true
	return instance
}

func convertExpensesFixed(expensesFixed []expensefixed.ExpenseFixed) []expense.Expense {
	var expenses []expense.Expense
	for _, entry := range expensesFixed {
		expenses = append(expenses, expense.Expense{time.Now(), entry.Description, entry.Amount})
	}
	return expenses
}
