package controllers_budget_test

import (
	"github.com/golang/mock/gomock"
	target "github.com/magleff/gobro/controllers/budget"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/mocks"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeInformation(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	currentBudget := &entities.Budget{}
	currentBudget.InitialBalance = decimal.NewFromFloat(1114.25)
	currentBudget.Expenses = []entities.Expense{
		entities.Expense{Amount: decimal.NewFromFloat(20.51), Checked: false},
		entities.Expense{Amount: decimal.NewFromFloat(-30.68), Checked: true},
		entities.Expense{Amount: decimal.NewFromFloat(10.05), Checked: false},
		entities.Expense{Amount: decimal.NewFromFloat(-18.36), Checked: false},
	}

	budgetController := mocks.NewMockBudgetController(mockCtrl)
	budgetController.EXPECT().Current().Return(currentBudget, nil)

	// Act
	information, _ := target.ComputeInformationDo(budgetController)

	// Assert
	assert.Equal(t, decimal.NewFromFloat(-49.04), information.TotalExpenses,
		"Should compute the total of expenses.")
	assert.Equal(t, decimal.NewFromFloat(30.56), information.TotalEarnings,
		"Should compute the total of earnings.")
	assert.Equal(t, decimal.NewFromFloat(-18.36), information.TotalUncheckedExpenses,
		"Should compute the total of unchecked expenses.")
	assert.Equal(t, decimal.NewFromFloat(1114.25), information.InitialBalance,
		"Should copy the initial balance.")
	assert.Equal(t, currentBudget.StartDate, information.StartDate,
		"Should copy the start date.")
	assert.Equal(t, decimal.NewFromFloat(-18.48), information.Difference,
		"Should compute the difference.")
	assert.Equal(t, decimal.NewFromFloat(1095.77), information.CurrentBalance,
		"Should compute the current balance.")
}
