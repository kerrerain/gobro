package cmd_test

import (
	cmdPackage "github.com/magleff/gobro/cmd"
	"github.com/spf13/cobra"
	"testing"
)

func TestRunInitCommand(t *testing.T) {
	// Arrange
	args := []string{"20.25"}

	budgetController := new(MockBudgetController)
	budgetController.On("CreateBudgetWithFixedExpenses", args[0]).Return(nil)

	command := cmdPackage.GobroInitCommand{}
	command.FlagPristine = false
	command.BudgetController = budgetController

	// Act
	command.Run(new(cobra.Command), args)

	// Assert
	budgetController.AssertExpectations(t)
}

func TestRunInitCommandDefaultInitialBalance(t *testing.T) {
	// Arrange
	args := []string{}

	budgetController := new(MockBudgetController)
	budgetController.On("CreateBudgetWithFixedExpenses", "0").Return(nil)

	command := cmdPackage.GobroInitCommand{}
	command.FlagPristine = false
	command.BudgetController = budgetController

	// Act
	command.Run(new(cobra.Command), args)

	// Assert
	budgetController.AssertExpectations(t)
}
