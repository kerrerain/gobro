package controllers_account_test

import (
	target "github.com/magleff/gobro/controllers/account"
	mocksModels "github.com/magleff/gobro/mocks/models"
	"github.com/magleff/gobro/models"
	"testing"
)

func TestList(t *testing.T) {
	// Arrange
	controller := target.Impl{}

	entity := mocksModels.Account{}
	entity.On("GetAll").Return([]models.Account{})

	// Act
	controller.List(entity)

	// Assert
	entity.AssertExpectations(t)
}
