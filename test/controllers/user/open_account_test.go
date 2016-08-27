package controllers_user_test

import (
	target "github.com/magleff/gobro/controllers/user"
	mocksModels "github.com/magleff/gobro/mocks/models"
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenAccount(t *testing.T) {
	// Arrange
	accountName := "main"
	userName := "default"

	userEntity := mocksModels.User{}
	accountEntity := mocksModels.Account{}

	accountEntity.On("FindByName", accountName).Return(&models.Account{})
	userEntity.On("FindByName", userName).Return(&models.User{})
	userEntity.On("Update", models.User{CurrentAccountName: accountName}).Return()

	// Act
	err := target.OpenDo(userEntity, accountEntity, userName, accountName)

	// Assert
	accountEntity.AssertExpectations(t)
	userEntity.AssertExpectations(t)
	assert.NoError(t, err, "")
}

func TestOpenAccountNoUser(t *testing.T) {
	// Arrange
	accountName := "main"
	userName := "default"

	userEntity := mocksModels.User{}
	accountEntity := mocksModels.Account{}

	accountEntity.On("FindByName", accountName).Return(&models.Account{})
	userEntity.On("FindByName", userName).Return(nil)

	// Act
	err := target.OpenDo(userEntity, accountEntity, userName, accountName)

	// Assert
	accountEntity.AssertExpectations(t)
	userEntity.AssertExpectations(t)
	assert.Error(t, err, "Should return an error if the user doesn't exist.")
}

func TestOpenAccountNoAccount(t *testing.T) {
	// Arrange
	accountName := "main"
	userName := "default"

	userEntity := mocksModels.User{}
	accountEntity := mocksModels.Account{}

	accountEntity.On("FindByName", accountName).Return(nil)

	// Act
	err := target.OpenDo(userEntity, accountEntity, userName, accountName)

	// Assert
	accountEntity.AssertExpectations(t)
	userEntity.AssertExpectations(t)
	assert.Error(t, err, "Should return an error if the account doesn't exist.")
}

func TestOpenAccountEmptyAccountName(t *testing.T) {
	// Arrange
	accountName := ""
	userName := "default"

	userEntity := mocksModels.User{}
	accountEntity := mocksModels.Account{}

	// Act
	err := target.OpenDo(userEntity, accountEntity, userName, accountName)

	// Assert
	assert.Error(t, err, "Should return an error if accountName is empty.")
}

func TestOpenAccountEmptyUserName(t *testing.T) {
	// Arrange
	accountName := "main"
	userName := ""

	userEntity := mocksModels.User{}
	accountEntity := mocksModels.Account{}

	// Act
	err := target.OpenDo(userEntity, accountEntity, userName, accountName)

	// Assert
	assert.Error(t, err, "Should return an error if userName is empty.")
}
