package controllers_budget_test

import (
	"github.com/golang/mock/gomock"
	target "github.com/magleff/gobro/controllers/budget"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/mocks"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestCurrent(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	currentUser := &entities.User{CurrentBudgetId: bson.NewObjectId()}

	budgetDao := mocks.NewMockBudgetDao(mockCtrl)
	budgetDao.EXPECT().FindById(currentUser.CurrentBudgetId).Return(&entities.Budget{}, nil)

	// Act
	_, _ = target.CurrentDo(budgetDao, currentUser)

	// Assert
}
