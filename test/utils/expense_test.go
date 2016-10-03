package utils_test

import (
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/utils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeTotalEarnings(t *testing.T) {
	// Arrange
	expenses := []entities.Expense{
		entities.Expense{Amount: decimal.NewFromFloat(20.51)},
		entities.Expense{Amount: decimal.NewFromFloat(-30.68)},
		entities.Expense{Amount: decimal.NewFromFloat(10.05)},
	}

	// Act
	totalEarnings := utils.ComputeTotalEarnings(expenses)

	// Assert
	assert.Equal(t, decimal.NewFromFloat(30.56), totalEarnings, "Should compute the total earnings (>0).")
}

func TestComputeTotalExpenses(t *testing.T) {
	// Arrange
	expenses := []entities.Expense{
		entities.Expense{Amount: decimal.NewFromFloat(-20.51)},
		entities.Expense{Amount: decimal.NewFromFloat(30.68)},
		entities.Expense{Amount: decimal.NewFromFloat(-10.05)},
	}

	// Act
	totalExpenses := utils.ComputeTotalExpenses(expenses)

	// Assert
	assert.Equal(t, decimal.NewFromFloat(-30.56), totalExpenses, "Should compute the total expenses (<=0).")
}

func TestComputeTotalUncheckedExpenses(t *testing.T) {
	// Arrange
	expenses := []entities.Expense{
		entities.Expense{Amount: decimal.NewFromFloat(-20.51), Checked: true},
		entities.Expense{Amount: decimal.NewFromFloat(-30.68), Checked: false},
		entities.Expense{Amount: decimal.NewFromFloat(-10.05), Checked: false},
	}

	// Act
	totalUnchekedExpenses := utils.ComputeTotalUncheckedExpenses(expenses)

	// Assert
	assert.Equal(t, decimal.NewFromFloat(-40.73),
		totalUnchekedExpenses, "Should compute the total unchecked expenses (<=0).")
}

// func TestFilterExpenses(t *testing.T) {
// 	checkedExpense := expense.NewExpense("50", "test")
// 	checkedExpense.Checked = true
// 	expenses := []expense.Expense{*checkedExpense,
// 		*expense.NewExpense("50", "test")}
// 	filteredExpenses := collections.Filter(expenses, func(expense expense.Expense) bool {
// 		return expense.Checked
// 	})
// 	if len(filteredExpenses) != 1 {
// 		t.Error("Expected filtered expenses to be size 1")
// 	}
// }
