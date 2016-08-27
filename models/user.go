package models

import (
	"github.com/magleff/gobro/database"
	"gopkg.in/mgo.v2/bson"
)

type UserEntity interface {
	FindByName(name string) *User
	Update(user User)
}

type User struct {
	ID                 bson.ObjectId `bson:"_id,omitempty"`
	CurrentBudgetId    bson.ObjectId `bson:"_id,omitempty"`
	CurrentAccountName string
	Name               string
}

func (e User) FindByName(name string) *User {
	var user User

	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("user").Find(bson.M{"name": name}).One(&user)
	})

	return &user
}

func (e User) Update(user User) {
	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("user").UpdateId(user.ID, user)
	})
}
