package cmd_test

import (
	"github.com/golang/mock/gomock"
	"github.com/magleff/gobro/cmd"
	"github.com/magleff/gobro/dto"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

// --- SETUP ---

type StatusCmdTestSuite struct {
	suite.Suite
	MockController       *gomock.Controller
	MockAccountDao       *mocks.MockAccountDao
	MockBudgetController *mocks.MockBudgetController
}

func (suite *StatusCmdTestSuite) SetupTest() {
	suite.MockController = gomock.NewController(suite.T())
	suite.MockAccountDao = mocks.NewMockAccountDao(suite.MockController)
	suite.MockBudgetController = mocks.NewMockBudgetController(suite.MockController)
}

func (suite *StatusCmdTestSuite) TearDownTest() {
	suite.MockController.Finish()
}

func TestStatusCmdTestSuite(t *testing.T) {
	suite.Run(t, new(StatusCmdTestSuite))
}

// --- TESTS ---

func (suite *StatusCmdTestSuite) TestStatusCmd() {
	// Arrange
	user := &entities.User{
		ID: bson.NewObjectId(),
		CliParams: entities.CliParams{
			CurrentAccountId: bson.NewObjectId(),
		},
	}

	suite.MockAccountDao.EXPECT().
		FindById(user.CliParams.CurrentAccountId).Return(&entities.Account{}, nil)

	suite.MockBudgetController.EXPECT().
		ComputeInformation(user.CliParams.CurrentAccountId).Return(&dto.BudgetInformation{}, nil)

	// Act
	err := cmd.StatusCmdDo([]string{}, suite.MockAccountDao, suite.MockBudgetController, user)

	// Assert
	assert.NoError(suite.T(), err, "Should not throw an error.")
}
