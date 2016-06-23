package budget

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/expensefixed"
	"gopkg.in/mgo.v2/bson"
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

func (self BudgetDatastore) GetCurrentBudget() Budget {
	session := self.DB.Session()
	var budgetSheets []Budget
	self.DB.Collection(session, "budget").Find(bson.M{"active": true}).All(&budgetSheets)
	defer session.Close()
	return budgetSheets[0]
}

func (self BudgetDatastore) Save(budget Budget) {
	session := self.DB.Session()
	self.DB.Collection(session, "budget").UpdateId(budget.ID, budget)
	defer session.Close()
}
