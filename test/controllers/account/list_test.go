package controllers_account_test

import (
	target "github.com/magleff/gobro/controllers/account"
	mocksModels "github.com/magleff/gobro/mocks/models"
	"github.com/magleff/gobro/models"
	"testing"
)

func TestList(t *testing.T) {
	// Arrange
	entity := mocksModels.Account{}
	entity.On("GetAll").Return([]models.Account{})

	// Act
	target.ListDo(entity)

	// Assert
	entity.AssertExpectations(t)
}
