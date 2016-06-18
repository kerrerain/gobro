package expensefixed

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ExpenseFixedDataStore struct {
	session *mgo.Session
}

func dataStore(session *mgo.Session) *ExpenseFixedDataStore {
	return &ExpenseFixedDataStore{session.Copy()}
}

func (eds ExpenseFixedDataStore) CreateExpenseFixed(amount float32, description string) {
	expensesFixed := eds.session.DB("").C("expenses-fixed")
	expensesFixed.Insert(ExpenseFixed{time.Now(), description, amount})
}

func (eds ExpenseFixedDataStore) ListExpensesFixed() []ExpenseFixed {
	var results []ExpenseFixed
	expensesFixed := eds.session.DB("").C("expenses-fixed")
	expensesFixed.Find(bson.M{}).All(&results)
	return results
}
