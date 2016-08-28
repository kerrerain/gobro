package controllers_user_test

import (
	"errors"
	target "github.com/magleff/gobro/controllers/user"
	mocksModels "github.com/magleff/gobro/mocks/models"
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenAccount(t *testing.T) {
	// Arrange
	accountName := "main"
	user := &models.User{}

	userEntity := mocksModels.User{}
	accountEntity := mocksModels.Account{}

	accountEntity.On("FindByName", accountName).Return(&models.Account{}, nil)
	userEntity.On("UpdateAccount", *user, models.Account{}).Return()

	// Act
	err := target.OpenAccountDo(userEntity, accountEntity, user, accountName)

	// Assert
	accountEntity.AssertExpectations(t)
	userEntity.AssertExpectations(t)
	assert.NoError(t, err, "")
}

func TestOpenAccountNoAccount(t *testing.T) {
	// Arrange
	accountName := "main"
	user := &models.User{}

	userEntity := mocksModels.User{}
	accountEntity := mocksModels.Account{}

	accountEntity.On("FindByName", accountName).Return(nil, errors.New("Doesn't exist."))

	// Act
	err := target.OpenAccountDo(userEntity, accountEntity, user, accountName)

	// Assert
	accountEntity.AssertExpectations(t)
	userEntity.AssertExpectations(t)
	assert.Error(t, err, "Should return an error if the account doesn't exist.")
}

func TestOpenAccountEmptyAccountName(t *testing.T) {
	// Arrange
	accountName := ""
	user := &models.User{}

	userEntity := mocksModels.User{}
	accountEntity := mocksModels.Account{}

	// Act
	err := target.OpenAccountDo(userEntity, accountEntity, user, accountName)

	// Assert
	assert.Error(t, err, "Should return an error if accountName is empty.")
}
