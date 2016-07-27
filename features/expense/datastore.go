package expense

import (
	"github.com/magleff/gobro/database"
)

type ExpenseDatastore struct {
	database.Datastore
}

func (self ExpenseDatastore) ImportExpensesIntoDB(entries []Expense) {
	self.ExecuteInSession(func() {
		expenses := self.Collection("expenses")
		for _, Expense := range entries {
			expenses.Insert(Expense)
		}
	})
}
