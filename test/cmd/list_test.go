package cmd_test

import (
	"github.com/magleff/gobro/cmd"
	mocksControllers "github.com/magleff/gobro/mocks/controllers"
	"github.com/magleff/gobro/models"
	"testing"
)

func TestListCmd(t *testing.T) {
	// Arrange
	controller := mocksControllers.Account{}
	controller.On("List").Return([]models.Account{})

	// Act
	cmd.ListCmdDo([]string{}, controller)

	// Assert
	controller.AssertExpectations(t)
}
