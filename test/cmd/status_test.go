package cmd_test

// import (
// 	"github.com/magleff/gobro/cmd"
// 	"github.com/magleff/gobro/models"
// 	"github.com/stretchr/testify/assert"
// 	"testing"
// )

// func TestStatusCmd(t *testing.T) {
// 	// Arrange
// 	budgetEntity := new(models.BudgetEntityMock)
// 	budgetEntity.On("GetCurrent").Return(&models.Budget{})

// 	// Act
// 	cmd.StatusCmd([]string{}, budgetEntity)

// 	// Assert
// 	budgetEntity.AssertExpectations(t)
// }

// func TestStatusCmdException(t *testing.T) {
// 	// Arrange
// 	budgetEntity := new(models.BudgetEntityMock)
// 	budgetEntity.On("GetCurrent").Return(nil)

// 	// Act
// 	err := cmd.StatusCmd([]string{}, budgetEntity)

// 	// Assert
// 	budgetEntity.AssertExpectations(t)
// 	assert.Error(t, err, "Should throw an error if there is not a budget")
// }
