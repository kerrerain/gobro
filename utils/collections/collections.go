package collections

import (
	"github.com/magleff/gobro/features/expense"
)

// As seen in https://gobyexample.com/collection-functions
// TODO make it generic? Find a battle-tested library?
func Filter(vs []expense.Expense, f func(expense.Expense) bool) []expense.Expense {
	vsf := make([]expense.Expense, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
