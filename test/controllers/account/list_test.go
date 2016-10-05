package controllers_account_test

import (
	"github.com/golang/mock/gomock"
	target "github.com/magleff/gobro/controllers/account"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/mocks"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

// --- SETUP ---

type ListTestSuite struct {
	suite.Suite
	MockAccountDao *mocks.MockAccountDao
	MockController *gomock.Controller
}

func (suite *ListTestSuite) SetupTest() {
	suite.MockController = gomock.NewController(suite.T())
	suite.MockAccountDao = mocks.NewMockAccountDao(suite.MockController)
}

func (suite *ListTestSuite) TearDownTest() {
	suite.MockController.Finish()
}

func TestListTestSuite(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}

// --- TESTS ---

func (suite *ListTestSuite) TestList() {
	// Arrange
	user := &entities.User{
		ID: bson.NewObjectId(),
	}

	suite.MockAccountDao.EXPECT().GetAll(user.ID).Return([]entities.Account{}, nil)

	// Act
	target.ListDo(suite.MockAccountDao, user.ID)

	// Assert
}
