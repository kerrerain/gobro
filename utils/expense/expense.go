package expense

import (
	"github.com/magleff/gobro/features/expense"
	"github.com/magleff/gobro/utils/collections"
	"github.com/shopspring/decimal"
)

func ComputeTotalEarnings(expenses []expense.Expense) decimal.Decimal {
	totalEarnings := decimal.NewFromFloat(0.00)
	for _, entry := range expenses {
		if entry.Amount.Cmp(decimal.NewFromFloat(0)) > 0 {
			totalEarnings = totalEarnings.Add(entry.Amount)
		}
	}
	return totalEarnings
}

func ComputeTotalExpenses(expenses []expense.Expense) decimal.Decimal {
	totalExpenses := decimal.NewFromFloat(0.00)
	for _, entry := range expenses {
		if entry.Amount.Cmp(decimal.NewFromFloat(0)) <= 0 {
			totalExpenses = totalExpenses.Add(entry.Amount)
		}
	}
	return totalExpenses
}

func ComputeTotalUncheckedExpenses(expenses []expense.Expense) decimal.Decimal {
	totalUncheckedExpenses := decimal.NewFromFloat(0.00)
	filteredExpenses := collections.Filter(expenses, func(obj expense.Expense) bool {
		return !obj.Checked
	})
	for _, entry := range filteredExpenses {
		if entry.Amount.Cmp(decimal.NewFromFloat(0)) <= 0 {
			totalUncheckedExpenses = totalUncheckedExpenses.Add(entry.Amount)
		}
	}
	return totalUncheckedExpenses
}
