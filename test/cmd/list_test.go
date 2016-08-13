package cmd_test

import (
	cmdPackage "github.com/magleff/gobro/cmd"
	accountPackage "github.com/magleff/gobro/features/account"
	budgetPackage "github.com/magleff/gobro/features/budget"
	"github.com/magleff/gobro/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ListCmdTestSuite struct {
	suite.Suite
	AccountController *mocks.MockAccountController
	BudgetController  *mocks.MockBudgetController
	Command           *cmdPackage.GobroListCommand
}

func (suite *ListCmdTestSuite) SetupTest() {
	suite.AccountController = new(mocks.MockAccountController)
	suite.BudgetController = new(mocks.MockBudgetController)

	suite.Command = new(cmdPackage.GobroListCommand)
	suite.Command.AccountController = suite.AccountController
	suite.Command.BudgetController = suite.BudgetController
}

func (suite *ListCmdTestSuite) TestListDefault() {
	// Arrange
	args := []string{}
	suite.BudgetController.On("CurrentBudget").Return(&budgetPackage.Budget{})

	// Act
	err := suite.Command.Run(new(cobra.Command), args)

	// Assert
	suite.BudgetController.AssertExpectations(suite.T())
	assert.NoError(suite.T(), err, "Should not throw an error.")
}

func (suite *ListCmdTestSuite) TestListAccount() {
	// Arrange
	args := []string{"account"}
	suite.AccountController.On("List").Return([]accountPackage.Account{})

	// Act
	err := suite.Command.Run(new(cobra.Command), args)

	// Assert
	suite.BudgetController.AssertExpectations(suite.T())
	assert.NoError(suite.T(), err, "Should not throw an error.")
}

func TestListCmdTestSuite(t *testing.T) {
	suite.Run(t, new(ListCmdTestSuite))
}
