package budget

import (
	"github.com/magleff/gobro/dto"
	"github.com/magleff/gobro/models"
)

type Controller interface {
	ComputeInformation() (*dto.BudgetInformation, error)
	Current() (*models.Budget, error)
}

type Impl struct{}
