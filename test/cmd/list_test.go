package cmd_test

import (
	"github.com/magleff/gobro/cmd"
	mocksControllers "github.com/magleff/gobro/mocks/controllers"
	mocksModels "github.com/magleff/gobro/mocks/models"
	"github.com/magleff/gobro/models"
	"testing"
)

func TestListCmd(t *testing.T) {
	// Arrange
	entity := mocksModels.Account{}

	controller := mocksControllers.Account{}
	controller.On("List", entity).Return([]models.Account{})

	// Act
	cmd.ListCmd([]string{}, controller, entity)

	// Assert
	controller.AssertExpectations(t)
}
