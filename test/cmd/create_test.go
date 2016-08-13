package cmd_test

import (
	cmdPackage "github.com/magleff/gobro/cmd"
	"github.com/magleff/gobro/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CreateCmdTestSuite struct {
	suite.Suite
	AccountController *mocks.MockAccountController
	Command           *cmdPackage.GobroCreateCommand
}

func (suite *CreateCmdTestSuite) SetupTest() {
	suite.AccountController = new(mocks.MockAccountController)
	suite.Command = new(cmdPackage.GobroCreateCommand)
	suite.Command.AccountController = suite.AccountController
}

func (suite *CreateCmdTestSuite) TestCreateWithEmptyType() {
	// Arrange
	args := []string{}

	// Act
	err := suite.Command.Run(new(cobra.Command), args)

	// Assert
	assert.Error(suite.T(), err, "Should throw an error if the type of element to create is not given.")
}

func (suite *CreateCmdTestSuite) TestCreateWithInvalidType() {
	// Arrange
	args := []string{"bazevhuxd"}

	// Act
	err := suite.Command.Run(new(cobra.Command), args)

	// Assert
	assert.Error(suite.T(), err, "Should throw an error if the type of element to create is not given.")
}

func (suite *CreateCmdTestSuite) TestCreateAccount() {
	// Arrange
	args := []string{"account", "Main account"}
	suite.AccountController.On("Create", args[1]).Return(nil)

	// Act
	err := suite.Command.Run(new(cobra.Command), args)

	// Assert
	suite.AccountController.AssertExpectations(suite.T())
	assert.NoError(suite.T(), err, "Should not throw an error.")
}

func (suite *CreateCmdTestSuite) TestCreateAccountWithEmptyName() {
	// Arrange
	args := []string{"account"}

	// Act
	err := suite.Command.Run(new(cobra.Command), args)

	// Assert
	assert.Error(suite.T(), err, "Should throw an error if the name of the account to create is not given.")
}

func TestCreateCmdTestSuite(t *testing.T) {
	suite.Run(t, new(CreateCmdTestSuite))
}
