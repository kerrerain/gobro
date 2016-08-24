package models

import (
	"github.com/magleff/gobro/database"
	"gopkg.in/mgo.v2/bson"
)

type AccountEntity interface {
	GetAll() []Account
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
