package dao

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/entities"
	"gopkg.in/mgo.v2/bson"
)

type AccountDao interface {
	GetAll() []entities.Account
	FindByName(string) (*entities.Account, error)
	Create(entities.User, entities.Account)
}

type AccountDaoImpl struct{}

func (e AccountDaoImpl) GetAll() []entities.Account {
	var accounts []entities.Account

	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("account").Find(bson.M{}).All(&accounts)
	})

	return accounts
}

func (e AccountDaoImpl) FindByName(name string) (*entities.Account, error) {
	var account entities.Account
	var err error

	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("account").Find(bson.M{"name": name}).One(&account)
	})

	return &account, err
}

func (e AccountDaoImpl) Create(user entities.User, account entities.Account) {
	account.UserId = user.ID
	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("account").Insert(account)
	})
}
