package budget

import (
	"github.com/magleff/gobro/dto"
	"gopkg.in/mgo.v2/bson"
)

type BudgetController interface {
	ComputeInformation(accountId bson.ObjectId) (*dto.BudgetInformation, error)
	Create(accountName string) error
}

type BudgetControllerImpl struct{}
