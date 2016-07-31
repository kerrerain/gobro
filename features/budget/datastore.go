package budget

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/expense"
	"github.com/shopspring/decimal"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type BudgetDatastore interface {
	CreateBudget(decimal.Decimal, []expense.Expense)
	CurrentBudget() *Budget
	Save(*Budget)
}

type BudgetDatastoreImpl struct {
	database.Datastore
}

func (self BudgetDatastoreImpl) CreateBudget(balance decimal.Decimal, expenses []expense.Expense) {
	self.ExecuteInSession(func() {
		self.Collection("budget").Insert(NewBudget(balance, expenses))
	})
}

// Gives the active budget sheet, or a nil value if currently there isn't one
func (self BudgetDatastoreImpl) CurrentBudget() *Budget {
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

func (self BudgetDatastoreImpl) Save(budget *Budget) {
	self.ExecuteInSession(func() {
		self.Collection("budget").UpdateId(budget.ID, *budget)
	})
}
