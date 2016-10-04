package dao

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/entities"
	"gopkg.in/mgo.v2/bson"
)

type UserDao interface {
	// Generic
	Create(user entities.User) error
	Update(user entities.User) error
	Delete(user entities.User) error
	FindById(userId bson.ObjectId) (*entities.User, error)
	// Specific
	FindByName(name string) (*entities.User, error)
}

type UserDaoImpl struct{}

func (e UserDaoImpl) Create(user entities.User) error {
	var err error
	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("user").Insert(user)
	})
	return err
}

func (e UserDaoImpl) Update(user entities.User) error {
	var err error
	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("user").UpdateId(user.ID, user)
	})
	return err
}

func (e UserDaoImpl) Delete(user entities.User) error {
	var err error
	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("user").RemoveId(user.ID)
	})
	return err
}

func (e UserDaoImpl) FindById(userId bson.ObjectId) (*entities.User, error) {
	var user entities.User
	var err error

	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("user").FindId(userId).One(&user)
	})

	return &user, err
}

func (e UserDaoImpl) FindByName(name string) (*entities.User, error) {
	var user entities.User
	var err error

	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("user").Find(bson.M{"name": name}).One(&user)
	})

	return &user, err
}
