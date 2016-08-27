package models

import (
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/mock"
)

type Budget struct {
	mock.Mock
}

func (m Budget) GetCurrent() *models.Budget {
	args := m.Called()
	if budget := args.Get(0); budget == nil {
		return nil
	} else {
		return budget.(*models.Budget)
	}
}
