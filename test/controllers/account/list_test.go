package controllers_account_test

import (
	"github.com/golang/mock/gomock"
	target "github.com/magleff/gobro/controllers/account"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/mocks"
	"testing"
)

func TestList(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	accountDao := mocks.NewMockAccountDao(mockCtrl)
	accountDao.EXPECT().GetAll().Return([]entities.Account{})

	// Act
	target.ListDo(accountDao)

	// Assert
}
