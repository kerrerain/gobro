package budget

import (
	"github.com/magleff/gobro/dto"
	"github.com/magleff/gobro/models"
)

type Controller interface {
	ComputeInformation(models.BudgetEntity) *dto.BudgetInformation
}

type Impl struct{}
