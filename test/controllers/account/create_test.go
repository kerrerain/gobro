package controllers_account_test

import (
	target "github.com/magleff/gobro/controllers/account"
	mocksModels "github.com/magleff/gobro/mocks/models"
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	// Arrange
	name := "main"
	controller := target.Impl{}

	entity := mocksModels.Account{}
	entity.On("FindByName", name).Return(nil)
	entity.On("Create", models.Account{Name: name}).Return()

	// Act
	err := controller.Create(entity, name)

	// Assert
	entity.AssertExpectations(t)
	assert.NoError(t, err, "Should not throw an error if there is not an account with the name.")
}

func TestCreateAlreadyExists(t *testing.T) {
	// Arrange
	name := "main"
	controller := target.Impl{}

	entity := mocksModels.Account{}
	entity.On("FindByName", name).Return(&models.Account{})

	// Act
	err := controller.Create(entity, name)

	// Assert
	entity.AssertExpectations(t)
	assert.Error(t, err, "Should throw an error if there is already an account with the name.")
}

func TestCreateEmptyName(t *testing.T) {
	// Arrange
	name := ""
	controller := target.Impl{}
	entity := mocksModels.Account{}

	// Act
	err := controller.Create(entity, name)

	// Assert
	entity.AssertExpectations(t)
	assert.Error(t, err, "Should throw an error if the name is empty.")
}
