package models

import (
	"github.com/stretchr/testify/mock"
)

type BudgetEntityMock struct {
	mock.Mock
}

func (m BudgetEntityMock) GetCurrent() *Budget {
	args := m.Called()
	if budget := args.Get(0); budget == nil {
		return nil
	} else {
		return budget.(*Budget)
	}
}
