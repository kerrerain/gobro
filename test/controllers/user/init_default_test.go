package controllers_user_test

import (
	"errors"
	"github.com/magleff/gobro/common"
	target "github.com/magleff/gobro/controllers/user"
	mocksModels "github.com/magleff/gobro/mocks/models"
	"github.com/magleff/gobro/models"
	"testing"
)

func TestInitDefault(t *testing.T) {
	// Arrange
	userName := common.DEFAULT_USER_NAME

	userEntity := mocksModels.User{}
	userEntity.On("FindByName", userName).Return(nil, errors.New("Doesn't exist."))
	userEntity.On("Create", models.User{Name: userName}).Return(nil)

	// Act
	target.InitDefaultDo(userEntity)

	// Assert
	userEntity.AssertExpectations(t)
}

func TestInitDefaultAlreadyExists(t *testing.T) {
	// Arrange
	userName := common.DEFAULT_USER_NAME

	userEntity := mocksModels.User{}
	userEntity.On("FindByName", userName).Return(&models.User{}, nil)

	// Act
	target.InitDefaultDo(userEntity)

	// Assert
	userEntity.AssertExpectations(t)
}
