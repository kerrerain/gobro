package models

import (
	"github.com/shopspring/decimal"
)

func ComputeTotalEarnings(expenses []Expense) decimal.Decimal {
	totalEarnings := decimal.NewFromFloat(0.00)
	for _, entry := range expenses {
		if entry.Amount.Cmp(decimal.NewFromFloat(0)) > 0 {
			totalEarnings = totalEarnings.Add(entry.Amount)
		}
	}
	return totalEarnings
}

func ComputeTotalExpenses(expenses []Expense) decimal.Decimal {
	totalExpenses := decimal.NewFromFloat(0.00)
	for _, entry := range expenses {
		if entry.Amount.Cmp(decimal.NewFromFloat(0)) <= 0 {
			totalExpenses = totalExpenses.Add(entry.Amount)
		}
	}
	return totalExpenses
}

func ComputeTotalUncheckedExpenses(expenses []Expense) decimal.Decimal {
	totalUncheckedExpenses := decimal.NewFromFloat(0.00)
	filteredExpenses := FilterExpenses(expenses, func(expense Expense) bool {
		return !expense.Checked
	})
	for _, entry := range filteredExpenses {
		if entry.Amount.Cmp(decimal.NewFromFloat(0)) <= 0 {
			totalUncheckedExpenses = totalUncheckedExpenses.Add(entry.Amount)
		}
	}
	return totalUncheckedExpenses
}

// As seen in https://gobyexample.com/collection-functions
func FilterExpenses(vs []Expense, f func(Expense) bool) []Expense {
	vsf := make([]Expense, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
