package dao

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/entities"
	"gopkg.in/mgo.v2/bson"
)

type AccountDao interface {
	// Generic
	Create(entities.Account)
	Update(account entities.Account) error
	Delete(account entities.Account) error
	FindById(accountId bson.ObjectId) (*entities.Account, error)
	// Specific
	GetAll(userId bson.ObjectId) ([]entities.Account, error)
	FindByName(userId bson.ObjectId, name string) (*entities.Account, error)
}

type AccountDaoImpl struct{}

func (e AccountDaoImpl) Create(account entities.Account) {
	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("account").Insert(account)
	})
}

func (e AccountDaoImpl) Update(account entities.Account) error {
	var err error
	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("account").UpdateId(account.ID, account)
	})
	return err
}

func (e AccountDaoImpl) Delete(account entities.Account) error {
	var err error
	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("account").RemoveId(account.ID)
	})
	return err
}

func (e AccountDaoImpl) FindById(accountId bson.ObjectId) (*entities.Account, error) {
	var account entities.Account
	var err error

	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("account").FindId(accountId).One(&account)
	})

	return &account, err
}

func (e AccountDaoImpl) GetAll(userId bson.ObjectId) ([]entities.Account, error) {
	var accounts []entities.Account
	var err error

	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("account").Find(bson.M{"userid": userId}).All(&accounts)
	})

	return accounts, err
}

func (e AccountDaoImpl) FindByName(userId bson.ObjectId, name string) (*entities.Account, error) {
	var account entities.Account
	var err error

	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("account").
			Find(bson.M{"name": name, "userid": userId}).One(&account)
	})

	return &account, err
}
