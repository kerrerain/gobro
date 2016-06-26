package budget

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/expense"
	"github.com/magleff/gobro/features/expensefixed"
	"log"
	"strings"
	"time"
)

type BudgetController struct {
	Datastore *BudgetDatastore
	DB        *database.Database
}

func NewController(DB *database.Database) *BudgetController {
	instance := new(BudgetController)
	instance.Datastore = NewDatastore(DB)
	instance.DB = DB
	return instance
}

func (self BudgetController) CreateBudget(balance string) {
	expensesFixed := expensefixed.NewDatastore(self.DB).ListExpensesFixed()
	if self.Datastore.CurrentBudget() == nil {
		self.Datastore.CreateBudget(expensesFixed, balance)
	} else {
		log.Fatal("There's already an active budget, use 'close' to close the current budget or 'rm' to remove it")
	}
}

func (self BudgetController) CreateBudgetWithoutExpensesFixed(balance string) {
	if self.Datastore.CurrentBudget() == nil {
		self.Datastore.CreateBudget([]expensefixed.ExpenseFixed{}, balance)
	} else {
		log.Fatal("There's already an active budget, use 'close' to close the current budget or 'rm' to remove it")
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

func (self BudgetController) CloseCurrentBudget() {
	currentBudget := self.Datastore.CurrentBudget()
	if currentBudget == nil {
		log.Fatal("There is not any active budget")
	} else {
		currentBudget.Active = false
		currentBudget.EndDate = time.Now()
		self.Datastore.Save(*currentBudget)
	}
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
