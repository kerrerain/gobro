package expensefixed

import (
	"github.com/magleff/gobro/database"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ExpenseFixedDataStore struct {
	DB *database.Database
}

func NewDatastore(DB *database.Database) *ExpenseFixedDataStore {
	instance := new(ExpenseFixedDataStore)
	instance.DB = DB
	return instance
}

func (self ExpenseFixedDataStore) CreateExpenseFixed(amount float32, description string) {
	session := self.DB.Session()
	self.DB.Collection(session, "expenses-fixed").Insert(ExpenseFixed{bson.NewObjectId(), time.Now(), description, amount})
	defer session.Close()
}

func (self ExpenseFixedDataStore) ListExpensesFixed() []ExpenseFixed {
	var results []ExpenseFixed
	session := self.DB.Session()
	self.DB.Collection(session, "expenses-fixed").Find(bson.M{}).All(&results)
	defer session.Close()
	return results
}

func (self ExpenseFixedDataStore) RemoveExpenseFixed(index int32) {
	expensesFixed := self.ListExpensesFixed()
	session := self.DB.Session()
	defer session.Close()
	self.DB.Collection(session, "expenses-fixed").Remove(bson.M{"_id": expensesFixed[index].ID})
}
