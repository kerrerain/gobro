package controllers_user_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/magleff/gobro/common"
	target "github.com/magleff/gobro/controllers/user"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/mocks"
	"testing"
)

func TestInitDefault(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userName := common.DEFAULT_USER_NAME

	userDao := mocks.NewMockUserDao(mockCtrl)
	userDao.EXPECT().FindByName(userName).Return(nil, errors.New("Doesn't exist."))
	userDao.EXPECT().Create(entities.User{Name: userName}).Return(nil)

	// Act
	target.InitDefaultDo(userDao)

	// Assert
}

func TestInitDefaultAlreadyExists(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userName := common.DEFAULT_USER_NAME

	userDao := mocks.NewMockUserDao(mockCtrl)
	userDao.EXPECT().FindByName(userName).Return(&entities.User{}, nil)

	// Act
	target.InitDefaultDo(userDao)

	// Assert
}
