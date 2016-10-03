package cmd_test

import (
	"github.com/golang/mock/gomock"
	"github.com/magleff/gobro/cmd"
	"github.com/magleff/gobro/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCmd(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	accountName := "main"

	controller := mocks.NewMockAccountController(mockCtrl)
	controller.EXPECT().Create(accountName).Return(nil)

	// Act
	err := cmd.CreateCmdDo([]string{accountName}, controller)

	// Assert
	assert.NoError(t, err, "Should not throw an error.")
}

func TestCreateCmdNoName(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	controller := mocks.NewMockAccountController(mockCtrl)

	// Act
	err := cmd.CreateCmdDo([]string{}, controller)

	// Assert
	assert.Error(t, err, "Should throw an error if the account name is not pro.")
}
