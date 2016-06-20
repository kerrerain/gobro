package budget

import (
	"github.com/magleff/gobro/expensefixed"
	"gopkg.in/mgo.v2"
)

type BudgetDatastore struct {
	session *mgo.Session
}

func DataStore(session *mgo.Session) *BudgetDatastore {
	return &BudgetDatastore{session.Copy()}
}

func collection(session *mgo.Session) *mgo.Collection {
	return session.DB("").C("budget")
}

func (eds BudgetDatastore) CreateBudget() {
	expensesFixed := expensefixed.DataStore(eds.session).ListExpensesFixed()
	collection(eds.session).Insert(NewBudget(expensesFixed))
}
