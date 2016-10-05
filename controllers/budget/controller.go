package budget

import (
	"github.com/magleff/gobro/dto"
	"gopkg.in/mgo.v2/bson"
)

type BudgetController interface {
	ComputeInformation(accountId bson.ObjectId) (*dto.BudgetInformation, error)
	Create(userId bson.ObjectId, accountId bson.ObjectId) error
}

type BudgetControllerImpl struct{}
