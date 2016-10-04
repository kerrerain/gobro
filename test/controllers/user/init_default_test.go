package controllers_user_test

import (
	"github.com/golang/mock/gomock"
	"github.com/magleff/gobro/common"
	target "github.com/magleff/gobro/controllers/user"
	"github.com/magleff/gobro/mocks"
	"testing"
)

func TestInitDefault(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userController := mocks.NewMockUserController(mockCtrl)
	userController.EXPECT().Create(common.DEFAULT_USER_NAME).Return(nil)

	// Act
	target.InitDefaultDo(userController)

	// Assert
}
