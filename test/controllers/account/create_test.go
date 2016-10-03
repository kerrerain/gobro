package controllers_account_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	target "github.com/magleff/gobro/controllers/account"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	name := "main"
	user := &entities.User{}

	accountDao := mocks.NewMockAccountDao(mockCtrl)
	accountDao.EXPECT().FindByName(name).Return(nil, errors.New("Doesn't exist."))
	accountDao.EXPECT().Create(*user, entities.Account{Name: name})

	// Act
	err := target.CreateDo(accountDao, user, name)

	// Assert
	assert.NoError(t, err, "Should not throw an error if there is not an account with the name.")
}

func TestCreateAlreadyExists(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	name := "main"
	user := &entities.User{}

	accountDao := mocks.NewMockAccountDao(mockCtrl)
	accountDao.EXPECT().FindByName(name).Return(&entities.Account{}, nil)

	// Act
	err := target.CreateDo(accountDao, user, name)

	// Assert
	assert.Error(t, err, "Should throw an error if there is already an account with the name.")
}

func TestCreateEmptyName(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	name := ""
	user := &entities.User{}
	accountDao := mocks.NewMockAccountDao(mockCtrl)

	// Act
	err := target.CreateDo(accountDao, user, name)

	// Assert
	assert.Error(t, err, "Should throw an error if the name is empty.")
}
