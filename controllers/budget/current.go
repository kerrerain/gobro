package budget

import (
	"errors"
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/session"
)

func (c BudgetControllerImpl) Current() (*entities.Budget, error) {
	// Manually inject entities
	return CurrentDo(dao.BudgetDaoImpl{}, session.GetCurrentUser())
}

func CurrentDo(budgetDao dao.BudgetDao, user *entities.User) (*entities.Budget, error) {
	currentBudget, err := budgetDao.FindById(user.CurrentBudgetId)

	if err != nil {
		return nil, errors.New("Test")
	}

	return currentBudget, nil
}
