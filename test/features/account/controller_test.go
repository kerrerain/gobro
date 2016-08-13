package account_test

import (
	accountPackage "github.com/magleff/gobro/features/account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccountControllerTestSuite struct {
	suite.Suite
	AccountDatastore *MockAccountDatastore
	Controller       *accountPackage.AccountControllerImpl
}

func (suite *AccountControllerTestSuite) SetupTest() {
	suite.AccountDatastore = new(MockAccountDatastore)
	suite.Controller = new(accountPackage.AccountControllerImpl)
	suite.Controller.AccountDatastore = suite.AccountDatastore
}

func (suite *AccountControllerTestSuite) TestCreate() {
	// Arrange
	name := "Main account"
	account := accountPackage.Account{Name: name}
	suite.AccountDatastore.On("Create", account).Return()

	// Act
	err := suite.Controller.Create(name)

	// Assert
	suite.AccountDatastore.AssertExpectations(suite.T())
	assert.NoError(suite.T(), err, "Should not throw an error.")
}

func (suite *AccountControllerTestSuite) TestList() {
	// Arrange
	suite.AccountDatastore.On("List").Return([]accountPackage.Account{})

	// Act
	suite.Controller.List()

	// Assert
	suite.AccountDatastore.AssertExpectations(suite.T())
}

func (suite *AccountControllerTestSuite) TestCurrent() {
	// Arrange
	suite.AccountDatastore.On("Current").Return(&accountPackage.Account{})

	// Act
	suite.Controller.Current()

	// Assert
	suite.AccountDatastore.AssertExpectations(suite.T())
}

func TestAccountControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AccountControllerTestSuite))
}
