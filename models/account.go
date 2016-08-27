package models

import (
	"github.com/magleff/gobro/database"
	"gopkg.in/mgo.v2/bson"
)

type AccountEntity interface {
	GetAll() []Account
	FindByName(string) *Account
	Create(Account)
}

type Account struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Name   string
	Active bool
}

func (e Account) GetAll() []Account {
	var accounts []Account

	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("account").Find(bson.M{}).All(&accounts)
	})

	return accounts
}

func (e Account) FindByName(name string) *Account {
	var account Account

	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("account").Find(bson.M{"name": name}).One(&account)
	})

	return &account
}

func (e Account) Create(account Account) {
	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("account").Insert(account)
	})
}
