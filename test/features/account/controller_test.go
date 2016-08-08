package account_test

import (
	accountPackage "github.com/magleff/gobro/features/account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccountControllerTestSuite struct {
	suite.Suite
	accountDatastore *MockAccountDatastore
	controller       *accountPackage.AccountControllerImpl
}

func (suite *AccountControllerTestSuite) SetupTest() {
	suite.accountDatastore = new(MockAccountDatastore)
	suite.controller = new(accountPackage.AccountControllerImpl)
	suite.controller.AccountDatastore = suite.accountDatastore
}

func (suite *AccountControllerTestSuite) TestCreate() {
	// Arrange
	name := "Main account"
	account := accountPackage.Account{Name: name}
	suite.accountDatastore.On("Create", account).Return()

	// Act
	err := suite.controller.Create(name)

	// Assert
	suite.accountDatastore.AssertExpectations(suite.T())
	assert.NoError(suite.T(), err, "Should not throw an error.")
}

func (suite *AccountControllerTestSuite) TestList() {
	// Arrange
	suite.accountDatastore.On("List").Return([]accountPackage.Account{})

	// Act
	suite.controller.List()

	// Assert
	suite.accountDatastore.AssertExpectations(suite.T())
}

func (suite *AccountControllerTestSuite) TestCurrent() {
	// Arrange
	suite.accountDatastore.On("Current").Return(&accountPackage.Account{})

	// Act
	suite.controller.Current()

	// Assert
	suite.accountDatastore.AssertExpectations(suite.T())
}

func TestAccountControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AccountControllerTestSuite))
}
