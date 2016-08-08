package account

import (
	"github.com/magleff/gobro/database"
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
	return nil
}
