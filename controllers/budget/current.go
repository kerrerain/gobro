package budget

import (
	"github.com/magleff/gobro/models"
	"github.com/magleff/gobro/session"
)

func (c Impl) Current() (*models.Budget, error) {
	// Manually inject entities
	return CurrentDo(models.Budget{}, session.GetCurrentUser())
}

func CurrentDo(budgetEntity models.BudgetEntity, user *models.User) (*models.Budget, error) {
	return nil, nil
}
