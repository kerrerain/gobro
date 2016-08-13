package account

import (
	"github.com/magleff/gobro/database"
	"gopkg.in/mgo.v2/bson"
)

type AccountDatastore interface {
	Create(Account)
	Current() *Account
	List() []Account
}

type AccountDatastoreImpl struct {
	database.Datastore
}

func (self AccountDatastoreImpl) Create(account Account) {
	self.ExecuteInSession(func() {
		self.Collection("account").Insert(account)
	})
}

func (self AccountDatastoreImpl) Current() *Account {
	return new(Account)
}

func (self AccountDatastoreImpl) List() []Account {
	var accounts []Account

	self.ExecuteInSession(func() {
		self.Collection("account").Find(bson.M{}).All(&accounts)
	})

	return accounts
}
