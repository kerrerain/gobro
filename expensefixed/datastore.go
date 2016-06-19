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

func collection(session *mgo.Session) *mgo.Collection {
	return session.DB("").C("expenses-fixed")
}

func (eds ExpenseFixedDataStore) CreateExpenseFixed(amount float32, description string) {
	collection(eds.session).Insert(ExpenseFixed{bson.NewObjectId(), time.Now(), description, amount})
}

func (eds ExpenseFixedDataStore) ListExpensesFixed() []ExpenseFixed {
	var results []ExpenseFixed
	collection(eds.session).Find(bson.M{}).All(&results)
	return results
}

func (eds ExpenseFixedDataStore) RemoveExpenseFixed(index int32) {
	expensesFixed := eds.ListExpensesFixed()
	collection(eds.session).Remove(bson.M{"_id": expensesFixed[index].ID})
}
