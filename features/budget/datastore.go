package budget

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/expensefixed"
	"gopkg.in/mgo.v2/bson"
	"log"
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

// Gives the active budget sheet, or a nil value if currently there isn't one
func (self BudgetDatastore) CurrentBudget() *Budget {
	var currentBudget *Budget

	session := self.DB.Session()
	var budgetSheets []Budget
	self.DB.Collection(session, "budget").Find(bson.M{"active": true}).All(&budgetSheets)

	defer session.Close()

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
	session := self.DB.Session()
	self.DB.Collection(session, "budget").UpdateId(budget.ID, budget)
	defer session.Close()
}
