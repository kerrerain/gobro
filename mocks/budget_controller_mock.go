package mocks

import (
	"github.com/magleff/gobro/features/budget"
	"github.com/magleff/gobro/features/expense"
	"github.com/stretchr/testify/mock"
)

type MockBudgetController struct {
	mock.Mock
}

func (m MockBudgetController) CreatePristineBudget(str string) error {
	args := m.Called(str)
	return args.Error(0)
}

func (m MockBudgetController) CreateBudgetWithFixedExpenses(str string) error {
	args := m.Called(str)
	return args.Error(0)
}

func (m MockBudgetController) CreateBudget(str string,
	expenses []expense.Expense) error {
	args := m.Called(str, expenses)
	return args.Error(0)
}

func (m MockBudgetController) SaveBudget(budget *budget.Budget) {
	m.Called(budget)
}

func (m MockBudgetController) CurrentBudget() *budget.Budget {
	args := m.Called()
	return args.Get(0).(*budget.Budget)
}

func (m MockBudgetController) AddExpenseToCurrentBudget(amount string,
	description string) error {
	args := m.Called(amount, description)
	return args.Error(0)
}

func (m MockBudgetController) AddRawExpensesToCurrentBudget(expenses []expense.Expense) error {
	args := m.Called(expenses)
	return args.Error(0)
}

func (m MockBudgetController) CloseCurrentBudget() error {
	args := m.Called()
	return args.Error(0)
}

func (m MockBudgetController) ComputeBudgetInfo() (*budget.BudgetInfo, error) {
	args := m.Called()
	return args.Get(0).(*budget.BudgetInfo), args.Error(1)
}
