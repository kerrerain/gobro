package models

import (
	"github.com/stretchr/testify/mock"
)

type AccountEntityMock struct {
	mock.Mock
}

func (m AccountEntityMock) GetAll() []Account {
	args := m.Called()
	if accounts := args.Get(0); accounts == nil {
		return nil
	} else {
		return accounts.([]Account)
	}
}
