package expensefixed

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/expense"
	"gopkg.in/mgo.v2/bson"
)

type ExpenseFixedDatastore struct {
	database.Datastore
}

func (self ExpenseFixedDatastore) CreateExpenseFixed(expenseFixed expense.Expense) {
	self.ExecuteInSession(func() {
		self.Collection("expenses-fixed").Insert(expenseFixed)
	})
}

func (self ExpenseFixedDatastore) ListExpensesFixed() []expense.Expense {
	var results []expense.Expense
	self.ExecuteInSession(func() {
		self.Collection("expenses-fixed").Find(bson.M{}).All(&results)
	})
	return results
}

func (self ExpenseFixedDatastore) RemoveExpenseFixed(index int32) {
	expensesFixed := self.ListExpensesFixed()
	self.ExecuteInSession(func() {
		self.Collection("expenses-fixed").Remove(bson.M{"_id": expensesFixed[index].ID})
	})
}
