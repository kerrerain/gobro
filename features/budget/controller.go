package budget

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/expense"
	"log"
	"strings"
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
	if self.Datastore.CurrentBudget() == nil {
		self.Datastore.CreateBudget()
	} else {
		log.Fatal("There's already an active budget, use 'close' to close the current budget")
	}
}

func (self BudgetController) CurrentBudget() *Budget {
	return self.Datastore.CurrentBudget()
}

func (self BudgetController) AddExpenseToCurrentBudget(amount string, description string) {
	currentBudget := self.Datastore.CurrentBudget()
	currentBudget.Expenses = append(currentBudget.Expenses, createExpense(amount, description))
	self.Datastore.Save(*currentBudget)
}

func createExpense(amount string, description string) expense.Expense {
	var newExpense expense.Expense
	if strings.Contains(amount, "+") {
		newExpense = *expense.NewResource(amount, description)
	} else {
		newExpense = *expense.NewExpense(amount, description)
	}
	return newExpense
}
