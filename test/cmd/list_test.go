package cmd_test

import (
	"github.com/magleff/gobro/cmd"
	"github.com/magleff/gobro/models"
	"testing"
)

func TestListCmd(t *testing.T) {
	// Arrange
	accountEntity := new(models.AccountEntityMock)
	accountEntity.On("GetAll").Return([]models.Account{})

	// Act
	cmd.ListCmd([]string{}, accountEntity)

	// Assert
	accountEntity.AssertExpectations(t)
}
