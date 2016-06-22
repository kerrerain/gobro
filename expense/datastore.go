package expense

import (
	"github.com/magleff/gobro/database"
)

type ExpenseDataStore struct {
	DB *database.Database
}

func NewDatastore(DB *database.Database) *ExpenseDataStore {
	instance := new(ExpenseDataStore)
	instance.DB = DB
	return instance
}

func (self ExpenseDataStore) ImportExpensesIntoDB(entries []Expense) {
	session := self.DB.Session()
	expenses := self.DB.Collection(session, "expenses")
	for _, Expense := range entries {
		expenses.Insert(Expense)
	}
	defer session.Close()
}
