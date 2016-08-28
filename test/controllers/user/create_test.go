package controllers_user_test

import (
	"errors"
	target "github.com/magleff/gobro/controllers/user"
	mocksModels "github.com/magleff/gobro/mocks/models"
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	// Arrange
	userName := "default"

	userEntity := mocksModels.User{}
	userEntity.On("FindByName", userName).Return(nil, errors.New("Doesn't exist."))
	userEntity.On("Create", models.User{Name: userName}).Return()

	// Act
	err := target.CreateDo(userEntity, userName)

	// Assert
	userEntity.AssertExpectations(t)
	assert.NoError(t, err, "")
}

func TestCreateAlreadyExists(t *testing.T) {
	// Arrange
	userName := "default"

	userEntity := mocksModels.User{}
	userEntity.On("FindByName", userName).Return(&models.User{}, nil)

	// Act
	err := target.CreateDo(userEntity, userName)

	// Assert
	userEntity.AssertExpectations(t)
	assert.Error(t, err, "Should return an error if the user already exists.")
}
