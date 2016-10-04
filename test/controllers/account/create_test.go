package controllers_account_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	target "github.com/magleff/gobro/controllers/account"
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
	MockAccountDao *mocks.MockAccountDao
	MockController *gomock.Controller
}

func (suite *CreateTestSuite) SetupTest() {
	suite.MockController = gomock.NewController(suite.T())
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
	name := "main"
	userId := bson.NewObjectId()

	suite.MockAccountDao.EXPECT().FindByName(userId, name).
		Return(nil, errors.New("Doesn't exist."))
	suite.MockAccountDao.EXPECT().Create(entities.Account{Name: name, UserId: userId})

	// Act
	err := target.CreateDo(suite.MockAccountDao, userId, name)

	// Assert
	assert.NoError(suite.T(), err,
		"Should not throw an error if there is not an account with the same name.")
}

func (suite *CreateTestSuite) TestCreateAlreadyExists() {
	// Arrange
	name := "main"
	userId := bson.NewObjectId()

	suite.MockAccountDao.EXPECT().FindByName(userId, name).
		Return(&entities.Account{}, nil)

	// Act
	err := target.CreateDo(suite.MockAccountDao, userId, name)

	// Assert
	assert.Error(suite.T(), err,
		"Should throw an error if there is already an account with the same name.")
}

func (suite *CreateTestSuite) TestCreateEmptyName() {
	// Arrange
	name := ""
	userId := bson.NewObjectId()

	// Act
	err := target.CreateDo(suite.MockAccountDao, userId, name)

	// Assert
	assert.Error(suite.T(), err, "Should throw an error if the name is empty.")
}
