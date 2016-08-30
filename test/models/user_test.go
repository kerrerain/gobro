package models

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
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
