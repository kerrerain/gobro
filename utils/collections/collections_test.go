package collections

import (
	"github.com/magleff/gobro/features/expense"
	"testing"
)

func TestFilter(t *testing.T) {
	checkedExpense := expense.NewExpense("50", "test")
	checkedExpense.Checked = true
	expenses := []expense.Expense{*checkedExpense,
		*expense.NewExpense("50", "test")}
	filteredExpenses := Filter(expenses, func(expense expense.Expense) bool {
		return expense.Checked
	})
	if len(filteredExpenses) != 1 {
		t.Error("Expected filtered expenses to be size 1")
	}
}
