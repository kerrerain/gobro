package cmd_test

import (
	"errors"
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

type OpenCmdTestSuite struct {
	suite.Suite
	MockUserDao    *mocks.MockUserDao
	MockAccountDao *mocks.MockAccountDao
	MockController *gomock.Controller
}

func (suite *OpenCmdTestSuite) SetupTest() {
	suite.MockController = gomock.NewController(suite.T())
	suite.MockAccountDao = mocks.NewMockAccountDao(suite.MockController)
	suite.MockUserDao = mocks.NewMockUserDao(suite.MockController)
}

func (suite *OpenCmdTestSuite) TearDownTest() {
	suite.MockController.Finish()
}

func TestOpenCmdTestSuite(t *testing.T) {
	suite.Run(t, new(OpenCmdTestSuite))
}

// --- TESTS ---

func (suite *OpenCmdTestSuite) TestOpenCmd() {
	// Arrange
	account := &entities.Account{
		ID:   bson.NewObjectId(),
		Name: "main",
	}

	user := &entities.User{
		ID: bson.NewObjectId(),
	}

	suite.MockAccountDao.EXPECT().FindByName(user.ID, account.Name).Return(account, nil)
	suite.MockUserDao.EXPECT().Update(entities.User{
		ID: user.ID,
		CliParams: entities.CliParams{
			CurrentAccountId: account.ID,
		},
	}).Return(nil)

	// Act
	err := cmd.OpenCmdDo([]string{account.Name}, suite.MockUserDao,
		suite.MockAccountDao, user)

	// Assert
	assert.NoError(suite.T(), err, "Should not throw an error.")
}

func (suite *OpenCmdTestSuite) TestOpenCmdNoAccountName() {
	// Arrange
	// Act
	err := cmd.OpenCmdDo([]string{}, suite.MockUserDao,
		suite.MockAccountDao, &entities.User{})

	// Assert
	assert.Error(suite.T(), err, "Should throw an error if the name is not provided.")
}

func (suite *OpenCmdTestSuite) TestOpenCmdAccountDoesNotExist() {
	// Arrange
	accountName := "main"

	user := &entities.User{
		ID: bson.NewObjectId(),
	}

	suite.MockAccountDao.EXPECT().
		FindByName(user.ID, accountName).Return(nil, errors.New("Not found."))

	// Act
	err := cmd.OpenCmdDo([]string{accountName}, suite.MockUserDao,
		suite.MockAccountDao, user)

	// Assert
	assert.Error(suite.T(), err, "Should throw an error if the account doesn't exist.")
}
