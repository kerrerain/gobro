package expensefixed

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/magleff/gobro/expense"
	"time"
)

type ExpenseFixedDataStore struct {
	session *mgo.Session
}

func dataStore(session *mgo.Session) *ExpenseFixedDataStore {
	return &ExpenseFixedDataStore{session.Copy()}
}

func (eds ExpenseFixedDataStore) CreateExpenseFixed(description string, amount float32) {
	expensesFixed := eds.session.DB("").C("expenses-fixed")
	expensesFixed.Insert(ExpenseFixed{expense.Expense{time.Now(), description, amount}})
}

func (eds ExpenseFixedDataStore) ListExpenses() []ExpenseFixed {
	var results []ExpenseFixed
	expensesFixed := eds.session.DB("").C("expenses-fixed")
	expensesFixed.Find(bson.M{}).All(&results)
	return results
}
