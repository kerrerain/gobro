package budget_test

import (
	budgetPackage "github.com/magleff/gobro/features/budget"
	"github.com/magleff/gobro/features/expense"
	"github.com/stretchr/testify/mock"
)

type BudgetDatastore interface {
	CreateBudget([]expense.Expense, string)
	CurrentBudget() *budgetPackage.Budget
	Save(*budgetPackage.Budget)
}

type MockBudgetDatastore struct {
	mock.Mock
}

func (m *MockBudgetDatastore) CreateBudget(expenses []expense.Expense, initialBalance string) {
	m.Called(expenses, initialBalance)
}

func (m *MockBudgetDatastore) CurrentBudget() *budgetPackage.Budget {
	args := m.Called()
	return args.Get(0).(*budgetPackage.Budget)
}

func (m *MockBudgetDatastore) Save(budget *budgetPackage.Budget) {
	m.Called(budget)
}
