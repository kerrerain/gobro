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

func TestCreate(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userName := "default"

	userDao := mocks.NewMockUserDao(mockCtrl)
	userDao.EXPECT().FindByName(userName).Return(nil, errors.New("Doesn't exist."))
	userDao.EXPECT().Create(entities.User{Name: userName}).Return(nil)

	// Act
	err := target.CreateDo(userDao, userName)

	// Assert
	assert.NoError(t, err, "")
}

func TestCreateAlreadyExists(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userName := "default"

	userDao := mocks.NewMockUserDao(mockCtrl)
	userDao.EXPECT().FindByName(userName).Return(&entities.User{}, nil)

	// Act
	err := target.CreateDo(userDao, userName)

	// Assert
	assert.Error(t, err, "Should return an error if the user already exists.")
}
