package budget_test

import (
	budgetPackage "github.com/magleff/gobro/features/budget"
	"github.com/magleff/gobro/features/expense"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewBudget(t *testing.T) {
	// Arrange
	initialExpenses := []expense.Expense{*expense.NewExpense("60.50", "test"),
		*expense.NewExpense("50.25", "test2")}
	balance := float32(32.52)

	// Act
	budget := budgetPackage.NewBudget(balance, initialExpenses)

	// Assert
	assert.Equal(t, float32(32.52), budget.InitialBalance, "Should init the budget with an initial balance.")
	assert.Equal(t, true, budget.Active, "Should make the budget active by default.")
	assert.Equal(t, time.Now().Format("2006-01-02"),
		budget.StartDate.Format("2006-01-02"),
		"Should set the start date to the current date.")
	assert.Equal(t, 2, len(budget.Expenses), "Should init the budget with expenses.")
}
