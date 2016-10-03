package models

import (
	"github.com/magleff/gobro/database"
	"gopkg.in/mgo.v2/bson"
)

type UserEntity interface {
	FindByName(name string) (*User, error)
	Update(user User) error
	Create(user User) error
}

type User struct {
	ID               bson.ObjectId `bson:"_id,omitempty"`
	CurrentAccountId bson.ObjectId `bson:"accountid,omitempty"`
	CurrentBudgetId  bson.ObjectId `bson:"budgetid,omitempty"`
	Name             string
}

func (e User) FindByName(name string) (*User, error) {
	var user User
	var err error

	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("user").Find(bson.M{"name": name}).One(&user)
	})

	return &user, err
}

func (e User) Update(user User) error {
	var err error
	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("user").UpdateId(user.ID, user)
	})
	return err
}

func (e User) Create(user User) error {
	var err error
	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("user").Insert(user)
	})
	return err
}
