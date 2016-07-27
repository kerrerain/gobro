package budget

import (
	"github.com/magleff/gobro/features/expense"
	"github.com/magleff/gobro/features/expensefixed"
	"log"
	"time"
)

type BudgetController struct {
	BudgetDatastore       *BudgetDatastore
	ExpenseFixedDatastore *expensefixed.ExpenseFixedDatastore
}

func NewBudgetController() *BudgetController {
	instance := new(BudgetController)
	instance.BudgetDatastore = new(BudgetDatastore)
	instance.ExpenseFixedDatastore = new(expensefixed.ExpenseFixedDatastore)
	return instance
}

func (self *BudgetController) CreateBudget(balance string) {
	expensesFixed := self.ExpenseFixedDatastore.ListExpensesFixed()

	if self.BudgetDatastore.CurrentBudget() == nil {
		self.BudgetDatastore.CreateBudget(expensesFixed, balance)
	} else {
		log.Fatal("There's already an active budget, use 'close' to close the current budget or 'rm' to remove it")
	}
}

func (self *BudgetController) SaveBudget(budget *Budget) {
	self.BudgetDatastore.Save(*budget)
}

func (self *BudgetController) CreateBudgetWithoutExpensesFixed(balance string) {
	if self.BudgetDatastore.CurrentBudget() == nil {
		self.BudgetDatastore.CreateBudget([]expense.Expense{}, balance)
	} else {
		log.Fatal("There's already an active budget, use 'close' to close the current budget or 'rm' to remove it")
	}
}

func (self *BudgetController) CurrentBudget() *Budget {
	return self.BudgetDatastore.CurrentBudget()
}

func (self *BudgetController) AddExpenseToCurrentBudget(amount string, description string) {
	currentBudget := self.BudgetDatastore.CurrentBudget()
	currentBudget.Expenses = append(currentBudget.Expenses, *expense.NewExpense(amount, description))
	self.BudgetDatastore.Save(*currentBudget)
}

func (self *BudgetController) AddRawExpensesToCurrentBudget(expenses []expense.Expense) {
	currentBudget := self.BudgetDatastore.CurrentBudget()
	for _, entry := range expenses {
		currentBudget.Expenses = append(currentBudget.Expenses, entry)
	}
	self.BudgetDatastore.Save(*currentBudget)
}

func (self *BudgetController) CloseCurrentBudget() {
	currentBudget := self.BudgetDatastore.CurrentBudget()
	if currentBudget == nil {
		log.Fatal("There is not any active budget")
	} else {
		currentBudget.Active = false
		currentBudget.EndDate = time.Now()
		self.BudgetDatastore.Save(*currentBudget)
	}
}
