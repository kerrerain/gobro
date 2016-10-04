package controllers_user_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	target "github.com/magleff/gobro/controllers/user"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// --- SETUP ---

type CreateTestSuite struct {
	suite.Suite
	MockUserDao    *mocks.MockUserDao
	MockController *gomock.Controller
}

func (suite *CreateTestSuite) SetupTest() {
	suite.MockController = gomock.NewController(suite.T())
	suite.MockUserDao = mocks.NewMockUserDao(suite.MockController)
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
	userName := "default"

	suite.MockUserDao.EXPECT().FindByName(userName).Return(nil, errors.New("Doesn't exist."))
	suite.MockUserDao.EXPECT().Create(entities.User{Name: userName}).Return(nil)

	// Act
	err := target.CreateDo(suite.MockUserDao, userName)

	// Assert
	assert.NoError(suite.T(), err, "")
}

func (suite *CreateTestSuite) TestCreateAlreadyExists() {
	// Arrange
	userName := "default"

	suite.MockUserDao.EXPECT().FindByName(userName).Return(&entities.User{}, nil)

	// Act
	err := target.CreateDo(suite.MockUserDao, userName)

	// Assert
	assert.Error(suite.T(), err, "Should return an error if the user already exists.")
}
