package cmd_test

import (
	cmdPackage "github.com/magleff/gobro/cmd"
	"github.com/magleff/gobro/mocks"
	"github.com/spf13/cobra"
	"testing"
)

func TestRunCloseCommand(t *testing.T) {
	// Arrange
	args := []string{}

	budgetController := new(mocks.MockBudgetController)
	budgetController.On("CloseCurrentBudget").Return(nil)

	command := cmdPackage.GobroCloseCommand{}
	command.BudgetController = budgetController

	// Act
	command.Run(new(cobra.Command), args)

	// Assert
	budgetController.AssertExpectations(t)
}
