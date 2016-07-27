package budget

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/expense"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type BudgetDatastore struct {
	database.Datastore
}

func (self BudgetDatastore) CreateBudget(expensesFixed []expense.Expense, balance string) {
	self.ExecuteInSession(func() {
		self.Collection("budget").Insert(NewBudgetWithExpensesFixed(expensesFixed, balance))
	})
}

// Gives the active budget sheet, or a nil value if currently there isn't one
func (self BudgetDatastore) CurrentBudget() *Budget {
	var currentBudget *Budget
	var budgetSheets []Budget

	self.ExecuteInSession(func() {
		self.Collection("budget").Find(bson.M{"active": true}).All(&budgetSheets)
	})

	if len(budgetSheets) > 1 {
		log.Fatal("There are more than one active budget!")
	} else if len(budgetSheets) == 1 {
		currentBudget = &budgetSheets[0]
	} else {
		currentBudget = nil
	}

	// Active budget sheet or nil
	return currentBudget
}

func (self BudgetDatastore) Save(budget Budget) {
	self.ExecuteInSession(func() {
		self.Collection("budget").UpdateId(budget.ID, budget)
	})
}
