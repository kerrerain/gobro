package budget_test

import (
	budgetPackage "github.com/magleff/gobro/features/budget"
	"github.com/magleff/gobro/features/expense"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/mock"
)

type MockBudgetDatastore struct {
	mock.Mock
}

func (m MockBudgetDatastore) CreateBudget(initialBalance decimal.Decimal,
	expenses []expense.Expense) {
	m.Called(initialBalance, expenses)
}

func (m MockBudgetDatastore) CurrentBudget() *budgetPackage.Budget {
	args := m.Called()
	if budget := args.Get(0); budget == nil {
		return nil
	} else {
		return budget.(*budgetPackage.Budget)
	}
}

func (m MockBudgetDatastore) Save(budget *budgetPackage.Budget) {
	m.Called(budget)
}
