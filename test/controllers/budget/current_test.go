package controllers_budget_test

import (
	target "github.com/magleff/gobro/controllers/budget"
	mocksModels "github.com/magleff/gobro/mocks/models"
	"github.com/magleff/gobro/models"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestCurrent(t *testing.T) {
	// Arrange
	currentUser := &models.User{CurrentAccountId: bson.NewObjectId()}

	entity := mocksModels.Budget{}
	entity.On("FindById", currentUser.CurrentAccountId).Return(&models.Budget{})

	// Act
	_, _ = target.CurrentDo(entity, currentUser)

	// Assert
	entity.AssertExpectations(t)
}
