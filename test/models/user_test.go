package models

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"log"
	"testing"
)

func TestCreate(t *testing.T) {
	// Arrange
	userName := "default"
	entity := models.User{}

	// Act
	err := entity.Create(models.User{Name: userName})

	// Assert
	var users []models.User

	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("user").Find(bson.M{}).All(&users)
	})

	assert.Equal(t, 1, len(users), "Should insert one user.")
	assert.Equal(t, users[0].Name, userName, "Should insert the right user.")
	assert.NoError(t, err, "Should not raise an error.")
}

func TestUpdate(t *testing.T) {
	// Arrange
	user := models.User{Name: "Baouh", ID: bson.NewObjectId()}
	entity := models.User{}

	database.ExecuteInSession(func(session database.Session) {
		if err := session.DefaultSchema().Collection("user").Insert(user); err != nil {
			log.Fatal(err)
		}
	})

	user.Name = "Changed"

	// Act
	err := entity.Update(user)

	// Assert
	var updatedUser models.User

	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("user").Find(bson.M{"_id": user.ID}).One(&updatedUser)
	})

	assert.Equal(t, "Changed", updatedUser.Name, "Should update the name of the user.")
	assert.NoError(t, err, "Should not raise an error.")
}
