package expense_test

import (
	"github.com/magleff/gobro/features/expense"
	expenseUtilsPackage "github.com/magleff/gobro/utils/expense"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeTotalEarnings(t *testing.T) {
	// Arrange
	expenses := []expense.Expense{
		expense.Expense{Amount: decimal.NewFromFloat(20.51)},
		expense.Expense{Amount: decimal.NewFromFloat(-30.68)},
		expense.Expense{Amount: decimal.NewFromFloat(10.05)},
	}

	// Act
	totalEarnings := expenseUtilsPackage.ComputeTotalEarnings(expenses)

	// Assert
	assert.Equal(t, decimal.NewFromFloat(30.56), totalEarnings, "Should compute the total earnings (>0).")
}

func TestComputeTotalExpenses(t *testing.T) {
	// Arrange
	expenses := []expense.Expense{
		expense.Expense{Amount: decimal.NewFromFloat(-20.51)},
		expense.Expense{Amount: decimal.NewFromFloat(30.68)},
		expense.Expense{Amount: decimal.NewFromFloat(-10.05)},
	}

	// Act
	totalExpenses := expenseUtilsPackage.ComputeTotalExpenses(expenses)

	// Assert
	assert.Equal(t, decimal.NewFromFloat(-30.56), totalExpenses, "Should compute the total expenses (<=0).")
}

func TestComputeTotalUncheckedExpenses(t *testing.T) {
	// Arrange
	expenses := []expense.Expense{
		expense.Expense{Amount: decimal.NewFromFloat(-20.51), Checked: true},
		expense.Expense{Amount: decimal.NewFromFloat(-30.68), Checked: false},
		expense.Expense{Amount: decimal.NewFromFloat(-10.05), Checked: false},
	}

	// Act
	totalUnchekedExpenses := expenseUtilsPackage.ComputeTotalUncheckedExpenses(expenses)

	// Assert
	assert.Equal(t, decimal.NewFromFloat(-40.73),
		totalUnchekedExpenses, "Should compute the total unchecked expenses (<=0).")
}
