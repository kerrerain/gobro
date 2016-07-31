package cmd_test

import (
	cmdPackage "github.com/magleff/gobro/cmd"
	"github.com/spf13/cobra"
	"testing"
)

func TestRunCloseCommand(t *testing.T) {
	// Arrange
	args := []string{}

	budgetController := new(MockBudgetController)
	budgetController.On("CloseCurrentBudget").Return(nil)

	command := cmdPackage.GobroCloseCommand{}
	command.BudgetController = budgetController

	// Act
	command.Run(new(cobra.Command), args)

	// Assert
	budgetController.AssertExpectations(t)
}
