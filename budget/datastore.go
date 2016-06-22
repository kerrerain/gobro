package budget

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/expensefixed"
)

type BudgetDatastore struct {
	DB *database.Database
}

func NewDatastore(DB *database.Database) *BudgetDatastore {
	instance := new(BudgetDatastore)
	instance.DB = DB
	return instance
}

func (self BudgetDatastore) CreateBudget() {
	session := self.DB.Session()
	expensesFixed := expensefixed.NewDatastore(self.DB).ListExpensesFixed()
	self.DB.Collection(session, "budget").Insert(NewBudget(expensesFixed))
	defer session.Close()
}
