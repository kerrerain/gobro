package cmd_test

import (
	"github.com/golang/mock/gomock"
	"github.com/magleff/gobro/cmd"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/mocks"
	"testing"
)

func TestListCmd(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	controller := mocks.NewMockAccountController(mockCtrl)
	controller.EXPECT().List().Return([]entities.Account{})

	// Act
	cmd.ListCmdDo([]string{}, controller)

	// Assert
}
