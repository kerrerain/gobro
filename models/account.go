package models

import (
	"github.com/magleff/gobro/database"
	"gopkg.in/mgo.v2/bson"
)

type AccountEntity interface {
	GetAll() []Account
	FindByName(string) (*Account, error)
	Create(User, Account)
}

type Account struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	UserId bson.ObjectId
	Name   string
	Label  string
	Active bool
}

func (e Account) GetAll() []Account {
	var accounts []Account

	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("account").Find(bson.M{}).All(&accounts)
	})

	return accounts
}

func (e Account) FindByName(name string) (*Account, error) {
	var account Account
	var err error

	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("account").Find(bson.M{"name": name}).One(&account)
	})

	return &account, err
}

func (e Account) Create(user User, account Account) {
	account.UserId = user.ID
	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("account").Insert(account)
	})
}
