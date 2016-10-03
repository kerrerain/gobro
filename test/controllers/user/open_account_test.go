package controllers_user_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	target "github.com/magleff/gobro/controllers/user"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenAccount(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	accountName := "main"
	user := &entities.User{}

	userDao := mocks.NewMockUserDao(mockCtrl)
	accountDao := mocks.NewMockAccountDao(mockCtrl)

	userDao.EXPECT().Update(*user).Return(nil)
	accountDao.EXPECT().FindByName(accountName).Return(&entities.Account{}, nil)

	// Act
	err := target.OpenAccountDo(userDao, accountDao, user, accountName)

	// Assert
	assert.NoError(t, err, "")
}

func TestOpenAccountNoAccount(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	accountName := "main"
	user := &entities.User{}

	userDao := mocks.NewMockUserDao(mockCtrl)
	accountDao := mocks.NewMockAccountDao(mockCtrl)

	accountDao.EXPECT().FindByName(accountName).Return(nil, errors.New("Doesn't exist."))

	// Act
	err := target.OpenAccountDo(userDao, accountDao, user, accountName)

	// Assert
	assert.Error(t, err, "Should return an error if the account doesn't exist.")
}

func TestOpenAccountEmptyAccountName(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	accountName := ""
	user := &entities.User{}

	userDao := mocks.NewMockUserDao(mockCtrl)
	accountDao := mocks.NewMockAccountDao(mockCtrl)

	// Act
	err := target.OpenAccountDo(userDao, accountDao, user, accountName)

	// Assert
	assert.Error(t, err, "Should return an error if accountName is empty.")
}
