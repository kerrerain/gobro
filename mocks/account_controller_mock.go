package mocks

import (
	"github.com/magleff/gobro/features/account"
	"github.com/stretchr/testify/mock"
)

type MockAccountController struct {
	mock.Mock
}

func (m MockAccountController) Create(str string) error {
	args := m.Called(str)
	return args.Error(0)
}

func (m MockAccountController) List() []account.Account {
	args := m.Called()
	return args.Get(0).([]account.Account)
}

func (m MockAccountController) Current() *account.Account {
	args := m.Called()
	return args.Get(0).(*account.Account)
}
