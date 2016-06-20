package expense

import (
	"gopkg.in/mgo.v2"
)

type ExpenseDataStore struct {
	session *mgo.Session
}

func DataStore(session *mgo.Session) *ExpenseDataStore {
	return &ExpenseDataStore{session.Copy()}
}

func (eds ExpenseDataStore) ImportExpensesIntoDB(entries []Expense) {
	expenses := eds.session.DB("").C("expenses")
	for _, Expense := range entries {
		expenses.Insert(Expense)
	}
}
