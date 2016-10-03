package models_test

import (
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/entities"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"log"
	"testing"
)

func TestCreate(t *testing.T) {
	// Arrange
	userName := "default"
	userDao := dao.UserDaoImpl{}

	// Act
	err := userDao.Create(entities.User{Name: userName})

	// Assert
	var users []entities.User

	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("user").Find(bson.M{}).All(&users)
	})

	assert.Equal(t, 1, len(users), "Should insert one user.")
	assert.Equal(t, users[0].Name, userName, "Should insert the right user.")
	assert.NoError(t, err, "Should not raise an error.")
}

func TestUpdate(t *testing.T) {
	// Arrange
	user := entities.User{Name: "Baouh", ID: bson.NewObjectId()}
	userDao := dao.UserDaoImpl{}

	database.ExecuteInSession(func(session database.Session) {
		if err := session.DefaultSchema().Collection("user").Insert(user); err != nil {
			log.Fatal(err)
		}
	})

	user.Name = "Changed"

	// Act
	err := userDao.Update(user)

	// Assert
	var updatedUser entities.User

	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("user").Find(bson.M{"_id": user.ID}).One(&updatedUser)
	})

	assert.Equal(t, "Changed", updatedUser.Name, "Should update the name of the user.")
	assert.NoError(t, err, "Should not raise an error.")
}
