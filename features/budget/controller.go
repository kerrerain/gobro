package budget

import (
	"errors"
	"github.com/magleff/gobro/features/expense"
	"github.com/magleff/gobro/features/expensefixed"
	"time"
)

type BudgetController interface {
	CreatePristineBudget(string) error
	CreateBudgetWithFixedExpenses(string) error
	CreateBudget(string, []expense.Expense) error
	SaveBudget(*Budget)
	CurrentBudget() *Budget
	AddExpenseToCurrentBudget(string, string) error
	AddRawExpensesToCurrentBudget([]expense.Expense)
	CloseCurrentBudget() error
}

type BudgetControllerImpl struct {
	BudgetDatastore       BudgetDatastore
	ExpenseFixedDatastore *expensefixed.ExpenseFixedDatastore
}

func NewBudgetController() BudgetController {
	instance := new(BudgetControllerImpl)
	instance.BudgetDatastore = new(BudgetDatastoreImpl)
	instance.ExpenseFixedDatastore = new(expensefixed.ExpenseFixedDatastore)
	return instance
}

func (self *BudgetControllerImpl) CreatePristineBudget(balance string) error {
	return self.CreateBudget(balance, []expense.Expense{})
}

func (self *BudgetControllerImpl) CreateBudgetWithFixedExpenses(balance string) error {
	expensesFixed := self.ExpenseFixedDatastore.ListExpensesFixed()
	return self.CreateBudget(balance, expensesFixed)
}

func (self *BudgetControllerImpl) CreateBudget(balance string, expenses []expense.Expense) error {
	if self.BudgetDatastore.CurrentBudget() == nil {
		self.BudgetDatastore.CreateBudget(balance, expenses)
	} else {
		return errors.New(`There's already an active budget,
			use 'close' to close the current budget or 'rm' to remove it`)
	}
	return nil
}

func (self *BudgetControllerImpl) SaveBudget(budget *Budget) {
	self.BudgetDatastore.Save(budget)
}

func (self *BudgetControllerImpl) CurrentBudget() *Budget {
	return self.BudgetDatastore.CurrentBudget()
}

func (self *BudgetControllerImpl) AddExpenseToCurrentBudget(amount string, description string) error {
	currentBudget := self.BudgetDatastore.CurrentBudget()
	if currentBudget != nil {
		currentBudget.Expenses = append(currentBudget.Expenses,
			*expense.NewExpense(amount, description))
		self.BudgetDatastore.Save(currentBudget)
	} else {
		return errors.New("There is not any active budget.")
	}
	return nil
}

func (self *BudgetControllerImpl) AddRawExpensesToCurrentBudget(expenses []expense.Expense) {
	currentBudget := self.BudgetDatastore.CurrentBudget()
	for _, entry := range expenses {
		currentBudget.Expenses = append(currentBudget.Expenses, entry)
	}
	self.BudgetDatastore.Save(currentBudget)
}

func (self *BudgetControllerImpl) CloseCurrentBudget() error {
	currentBudget := self.BudgetDatastore.CurrentBudget()
	if currentBudget != nil {
		currentBudget.Active = false
		currentBudget.EndDate = time.Now()
		self.BudgetDatastore.Save(currentBudget)
	} else {
		return errors.New("There is not any active budget.")
	}
	return nil
}
