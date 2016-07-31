package budget

import (
	"errors"
	"github.com/magleff/gobro/features/expense"
	"github.com/magleff/gobro/features/expensefixed"
	"github.com/magleff/gobro/utils/amount"
	"time"
)

type BudgetController interface {
	CreatePristineBudget(string) error
	CreateBudgetWithFixedExpenses(string) error
	CreateBudget(string, []expense.Expense) error
	SaveBudget(*Budget)
	CurrentBudget() *Budget
	AddExpenseToCurrentBudget(string, string) error
	AddRawExpensesToCurrentBudget([]expense.Expense) error
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

// Creates a budget with an initial balance (mandatory) and does not compute
// the initial expenses from a payment schedule.
//
// Returns an error if there is not any active budget for the moment.
// Returns and error if the initial balance is not set or invalid.
func (self BudgetControllerImpl) CreatePristineBudget(balance string) error {
	return self.CreateBudget(balance, []expense.Expense{})
}

// Creates a budget with an initial balance (mandatory) and computes
// the initial expenses from a payment schedule.
//
// Returns an error if there is not any active budget for the moment.
// Returns and error if the initial balance is not set or invalid.
func (self BudgetControllerImpl) CreateBudgetWithFixedExpenses(balance string) error {
	expensesFixed := self.ExpenseFixedDatastore.ListExpensesFixed()
	return self.CreateBudget(balance, expensesFixed)
}

// Creates a new budget with an initial balance (mandatory), and sets the initial expenses.
//
// Returns an error if there is not any active budget for the moment.
// Returns and error if the initial balance is not set or invalid.
func (self BudgetControllerImpl) CreateBudget(balance string, expenses []expense.Expense) error {
	parsedBalance, err := amount.ParseString(balance)

	if err != nil {
		return err
	}

	if self.BudgetDatastore.CurrentBudget() == nil {
		self.BudgetDatastore.CreateBudget(parsedBalance, expenses)
	} else {
		return errors.New("There's already an active budget, " +
			"use 'close' to close the current budget or 'rm' to remove it.")
	}

	return nil
}

func (self BudgetControllerImpl) SaveBudget(budget *Budget) {
	self.BudgetDatastore.Save(budget)
}

func (self BudgetControllerImpl) CurrentBudget() *Budget {
	return self.BudgetDatastore.CurrentBudget()
}

func (self BudgetControllerImpl) AddExpenseToCurrentBudget(amount string, description string) error {
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

func (self BudgetControllerImpl) AddRawExpensesToCurrentBudget(expenses []expense.Expense) error {
	currentBudget := self.BudgetDatastore.CurrentBudget()
	if currentBudget != nil {
		for _, entry := range expenses {
			currentBudget.Expenses = append(currentBudget.Expenses, entry)
		}
		self.BudgetDatastore.Save(currentBudget)
	} else {
		return errors.New("There is not any active budget.")
	}
	return nil
}

func (self BudgetControllerImpl) CloseCurrentBudget() error {
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
