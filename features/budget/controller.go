package budget

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/expense"
)

type BudgetController struct {
	Datastore *BudgetDatastore
}

func NewController(DB *database.Database) *BudgetController {
	instance := new(BudgetController)
	instance.Datastore = NewDatastore(DB)
	return instance
}

func (self BudgetController) CreateBudget() {
	self.Datastore.CreateBudget()
}

func (self BudgetController) AddExpenseToCurrentBudget(amount string, description string) {
	currentBudget := self.Datastore.GetCurrentBudget()
	currentBudget.Expenses = append(currentBudget.Expenses, *expense.NewExpense(amount, description))
	self.Datastore.Save(currentBudget)
}
