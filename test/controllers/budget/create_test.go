package controllers_budget_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	target "github.com/magleff/gobro/controllers/budget"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

// --- SETUP ---

type CreateTestSuite struct {
	suite.Suite
	MockBudgetDao  *mocks.MockBudgetDao
	MockAccountDao *mocks.MockAccountDao
	MockController *gomock.Controller
}

func (suite *CreateTestSuite) SetupTest() {
	suite.MockController = gomock.NewController(suite.T())
	suite.MockBudgetDao = mocks.NewMockBudgetDao(suite.MockController)
	suite.MockAccountDao = mocks.NewMockAccountDao(suite.MockController)
}

func (suite *CreateTestSuite) TearDownTest() {
	suite.MockController.Finish()
}

func TestCreateTestSuite(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}

// --- TESTS ---

func (suite *CreateTestSuite) TestCreate() {
	// Arrange
	accountId := bson.NewObjectId()
	userId := bson.NewObjectId()

	suite.MockAccountDao.EXPECT().FindById(accountId).Return(&entities.Account{}, nil)
	suite.MockBudgetDao.EXPECT().FindActiveBudget(accountId).Return(nil, errors.New("Not found."))
	suite.MockBudgetDao.EXPECT().Create(entities.Budget{AccountId: accountId,
		UserId: userId}).Return(nil)

	// Act
	err := target.CreateDo(suite.MockBudgetDao, suite.MockAccountDao, userId, accountId)

	// Assert
	assert.NoError(suite.T(), err, "Should not throw an error.")
}

func (suite *CreateTestSuite) TestCreateAccountDoesNotExist() {
	// Arrange
	accountId := bson.NewObjectId()
	userId := bson.NewObjectId()

	suite.MockAccountDao.EXPECT().FindById(accountId).Return(nil, errors.New("Not found."))

	// Act
	err := target.CreateDo(suite.MockBudgetDao, suite.MockAccountDao, userId, accountId)

	// Assert
	assert.Error(suite.T(), err, "Should throw an error if the account does not exist.")
}

func (suite *CreateTestSuite) TestCreateAlreadyAnOpenedBudget() {
	// Arrange
	accountId := bson.NewObjectId()
	userId := bson.NewObjectId()

	suite.MockAccountDao.EXPECT().FindById(accountId).Return(&entities.Account{}, nil)
	suite.MockBudgetDao.EXPECT().FindActiveBudget(accountId).Return(&entities.Budget{}, nil)

	// Act
	err := target.CreateDo(suite.MockBudgetDao, suite.MockAccountDao, userId, accountId)

	// Assert
	assert.Error(suite.T(), err, "Should throw an error if there's already an active budget.")
}
