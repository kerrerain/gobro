package mocks

import (
	"github.com/magleff/gobro/features/account"
	"github.com/stretchr/testify/mock"
	"log"
)

type MockAccountController struct {
	mock.Mock
}

func (m MockAccountController) Create(str string) error {
	args := m.Called(str)
	log.Println(args)
	return args.Error(0)
}

func (m MockAccountController) List() []account.Account {
	m.Called()
	return nil
}

func (m MockAccountController) Current() *account.Account {
	args := m.Called()
	return args.Get(0).(*account.Account)
}
