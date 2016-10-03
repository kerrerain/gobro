package models

import (
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/mock"
	"gopkg.in/mgo.v2/bson"
)

type Budget struct {
	mock.Mock
}

func (m Budget) FindById(budgetId bson.ObjectId) (*models.Budget, error) {
	args := m.Called(budgetId)
	if budget := args.Get(0); budget == nil {
		return nil, args.Error(1)
	} else {
		return budget.(*models.Budget), args.Error(1)
	}
}
