package models

import (
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/mock"
)

type Account struct {
	mock.Mock
}

func (m Account) GetAll() []models.Account {
	args := m.Called()
	if accounts := args.Get(0); accounts == nil {
		return nil
	} else {
		return accounts.([]models.Account)
	}
}

func (m Account) FindByName(name string) (*models.Account, error) {
	args := m.Called(name)
	if account := args.Get(0); account == nil {
		return nil, args.Error(1)
	} else {
		return account.(*models.Account), args.Error(1)
	}
}

func (m Account) Create(user models.User, account models.Account) {
	m.Called(user, account)
}
