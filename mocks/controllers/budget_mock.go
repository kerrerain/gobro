package controllers

import (
	"github.com/magleff/gobro/dto"
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/mock"
)

type Budget struct {
	mock.Mock
}

func (m Budget) ComputeInformation() (*dto.BudgetInformation, error) {
	args := m.Called()
	if info := args.Get(0); info == nil {
		return nil, args.Error(1)
	} else {
		return info.(*dto.BudgetInformation), args.Error(1)
	}
}

func (m Budget) Current() (*models.Budget, error) {
	args := m.Called()
	if budget := args.Get(0); budget == nil {
		return nil, args.Error(1)
	} else {
		return budget.(*models.Budget), args.Error(1)
	}
}
