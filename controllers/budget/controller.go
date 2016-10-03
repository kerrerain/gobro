package budget

import (
	"github.com/magleff/gobro/dto"
	"github.com/magleff/gobro/entities"
)

type BudgetController interface {
	ComputeInformation() (*dto.BudgetInformation, error)
	Current() (*entities.Budget, error)
}

type BudgetControllerImpl struct{}
