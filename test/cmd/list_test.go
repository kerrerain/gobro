package cmd_test

import (
	"github.com/golang/mock/gomock"
	"github.com/magleff/gobro/cmd"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

// --- SETUP ---

type ListCmdTestSuite struct {
	suite.Suite
	MockAccountController *mocks.MockAccountController
	MockBudgetController  *mocks.MockBudgetController
	MockController        *gomock.Controller
}

func (suite *ListCmdTestSuite) SetupTest() {
	suite.MockController = gomock.NewController(suite.T())
	suite.MockAccountController = mocks.NewMockAccountController(suite.MockController)
}

func (suite *ListCmdTestSuite) TearDownTest() {
	suite.MockController.Finish()
}

func TestListCmdTestSuite(t *testing.T) {
	suite.Run(t, new(ListCmdTestSuite))
}

// --- TESTS ---

func (suite *ListCmdTestSuite) TestListCmd() {
	// Arrange
	user := &entities.User{
		ID: bson.NewObjectId(),
	}

	suite.MockAccountController.EXPECT().List(user.ID).Return([]entities.Account{}, nil)

	// Act
	err := cmd.ListCmdDo([]string{}, suite.MockAccountController, user)

	// Assert
	assert.NoError(suite.T(), err, "Should not throw an error.")
}
