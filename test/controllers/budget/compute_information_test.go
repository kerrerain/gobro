package controllers_budget_test

import (
	target "github.com/magleff/gobro/controllers/budget"
	mocksModels "github.com/magleff/gobro/mocks/models"
	"github.com/magleff/gobro/models"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeInformation(t *testing.T) {
	// Arrange
	controller := target.Impl{}

	currentBudget := &models.Budget{}
	currentBudget.InitialBalance = decimal.NewFromFloat(1114.25)
	currentBudget.Expenses = []models.Expense{
		models.Expense{Amount: decimal.NewFromFloat(20.51), Checked: false},
		models.Expense{Amount: decimal.NewFromFloat(-30.68), Checked: true},
		models.Expense{Amount: decimal.NewFromFloat(10.05), Checked: false},
		models.Expense{Amount: decimal.NewFromFloat(-18.36), Checked: false},
	}

	entity := mocksModels.Budget{}
	entity.On("GetCurrent").Return(currentBudget)

	// Act
	information := controller.ComputeInformation(entity)

	// Assert
	entity.AssertExpectations(t)

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
