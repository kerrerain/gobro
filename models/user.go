package models

import (
	"github.com/magleff/gobro/database"
	"gopkg.in/mgo.v2/bson"
)

type UserEntity interface {
	FindByName(name string) (*User, error)
	Update(user User)
	Create(user User) error
	UpdateAccount(user User, account Account)
}

type User struct {
	ID                bson.ObjectId `bson:"_id,omitempty"`
	CurrentAccountId  bson.ObjectId `bson:"accountid,omitempty"`
	CurrentBudgetsIds []bson.ObjectId
	Name              string
}

func (e User) FindByName(name string) (*User, error) {
	var user User
	var err error

	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("user").Find(bson.M{"name": name}).One(&user)
	})

	return &user, err
}

func (e User) Update(user User) {
	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("user").UpdateId(user.ID, user)
	})
}

func (e User) Create(user User) error {
	var err error
	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("user").Insert(user)
	})
	return err
}

func (e User) UpdateAccount(user User, account Account) {
	user.CurrentAccountId = account.ID
	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("user").UpdateId(user.ID, user)
	})
}
