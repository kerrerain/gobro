package cmd_test

import (
	cmdPackage "github.com/magleff/gobro/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunInitCommand(t *testing.T) {
	// Arrange
	args := []string{"20.25"}

	budgetController := new(MockBudgetController)
	budgetController.On("CreatePristineBudget", args[0]).Return(nil)

	command := cmdPackage.GobroInitCommand{}
	command.FlagFixed = false
	command.BudgetController = budgetController

	// Act
	command.Run(new(cobra.Command), args)

	// Assert
	budgetController.AssertExpectations(t)
}

func TestRunInitCommandDefaultInitialBalance(t *testing.T) {
	// Arrange
	args := []string{}

	command := cmdPackage.GobroInitCommand{}

	// Act
	err := command.Run(new(cobra.Command), args)

	// Assert
	assert.Error(t, err, "Should throw an error if the initial balance is not given.")
}

func TestRunInitCommandFlagFixed(t *testing.T) {
	// Arrange
	args := []string{"20.25"}

	budgetController := new(MockBudgetController)
	budgetController.On("CreateBudgetWithFixedExpenses", args[0]).Return(nil)

	command := cmdPackage.GobroInitCommand{}
	command.FlagFixed = true
	command.BudgetController = budgetController

	// Act
	command.Run(new(cobra.Command), args)

	// Assert
	budgetController.AssertExpectations(t)
}
