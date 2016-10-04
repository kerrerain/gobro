package controllers_budget_test

import (
	"github.com/golang/mock/gomock"
	target "github.com/magleff/gobro/controllers/budget"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/mocks"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

// --- SETUP ---

type ComputeInformationTestSuite struct {
	suite.Suite
	MockBudgetDao  *mocks.MockBudgetDao
	MockController *gomock.Controller
}

func (suite *ComputeInformationTestSuite) SetupTest() {
	suite.MockController = gomock.NewController(suite.T())
	suite.MockBudgetDao = mocks.NewMockBudgetDao(suite.MockController)
}

func (suite *ComputeInformationTestSuite) TearDownTest() {
	suite.MockController.Finish()
}

func TestComputeInformationTestSuite(t *testing.T) {
	suite.Run(t, new(ComputeInformationTestSuite))
}

// --- TESTS ---

func (suite *CreateTestSuite) TestComputeInformation() {
	// Arrange
	accountId := bson.NewObjectId()

	currentBudget := &entities.Budget{}
	currentBudget.InitialBalance = decimal.NewFromFloat(1114.25)
	currentBudget.Expenses = []entities.Expense{
		entities.Expense{Amount: decimal.NewFromFloat(20.51), Checked: false},
		entities.Expense{Amount: decimal.NewFromFloat(-30.68), Checked: true},
		entities.Expense{Amount: decimal.NewFromFloat(10.05), Checked: false},
		entities.Expense{Amount: decimal.NewFromFloat(-18.36), Checked: false},
	}

	suite.MockBudgetDao.EXPECT().FindActiveBudget(accountId).Return(currentBudget, nil)

	// Act
	information, _ := target.ComputeInformationDo(suite.MockBudgetDao, accountId)

	// Assert
	assert.Equal(suite.T(), decimal.NewFromFloat(-49.04), information.TotalExpenses,
		"Should compute the total of expenses.")
	assert.Equal(suite.T(), decimal.NewFromFloat(30.56), information.TotalEarnings,
		"Should compute the total of earnings.")
	assert.Equal(suite.T(), decimal.NewFromFloat(-18.36), information.TotalUncheckedExpenses,
		"Should compute the total of unchecked expenses.")
	assert.Equal(suite.T(), decimal.NewFromFloat(1114.25), information.InitialBalance,
		"Should copy the initial balance.")
	assert.Equal(suite.T(), currentBudget.StartDate, information.StartDate,
		"Should copy the start date.")
	assert.Equal(suite.T(), decimal.NewFromFloat(-18.48), information.Difference,
		"Should compute the difference.")
	assert.Equal(suite.T(), decimal.NewFromFloat(1095.77), information.CurrentBalance,
		"Should compute the current balance.")
}
