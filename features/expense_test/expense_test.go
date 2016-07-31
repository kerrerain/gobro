package expense_test

import (
	expensePackage "github.com/magleff/gobro/features/expense"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewExpense(t *testing.T) {
	// Arrange
	amount := "165.344"
	description := "Restaurant"

	// Act
	expense := expensePackage.NewExpense(amount, description)

	// Assert
	assert.NotNil(t, expense, "Should return an expense.")
	assert.Equal(t, description, expense.Description, "Should copy the description.")
	assert.Equal(t, decimal.NewFromFloat(-165.34), expense.Amount, "Should parse the amount.")
	assert.Equal(t, false, expense.Checked, "Should be unchecked by default.")
	assert.Equal(t, time.Now().Format("2006-01-02"),
		expense.Date.Format("2006-01-02"),
		"Should set the date to the current date.")
}
