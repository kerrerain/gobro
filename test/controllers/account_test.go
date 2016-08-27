package controllers_test

import (
	"github.com/magleff/gobro/controllers"
	mocksModels "github.com/magleff/gobro/mocks/models"
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	// Arrange
	controller := controllers.Account{}

	entity := mocksModels.Account{}
	entity.On("GetAll").Return([]models.Account{})

	// Act
	controller.List(entity)

	// Assert
	entity.AssertExpectations(t)
}

func TestOpen(t *testing.T) {
	// Arrange
	name := "main"
	controller := controllers.Account{}

	entity := mocksModels.Account{}
	entity.On("FindByName", name).Return(nil)
	entity.On("Create", models.Account{Name: name}).Return()

	// Act
	err := controller.Open(entity, name)

	// Assert
	entity.AssertExpectations(t)
	assert.NoError(t, err, "Should not throw an error if there is not an account with the name.")
}

func TestOpenAlreadyExists(t *testing.T) {
	// Arrange
	name := "main"
	controller := controllers.Account{}

	entity := mocksModels.Account{}
	entity.On("FindByName", name).Return(&models.Account{})

	// Act
	err := controller.Open(entity, name)

	// Assert
	entity.AssertExpectations(t)
	assert.Error(t, err, "Should throw an error if there is already an account with the name.")
}

func TestOpenEmptyName(t *testing.T) {
	// Arrange
	name := ""
	controller := controllers.Account{}
	entity := mocksModels.Account{}

	// Act
	err := controller.Open(entity, name)

	// Assert
	entity.AssertExpectations(t)
	assert.Error(t, err, "Should throw an error if the name is empty.")
}
