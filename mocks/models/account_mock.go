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

func (m Account) FindByName(name string) *models.Account {
	args := m.Called(name)
	if account := args.Get(0); account == nil {
		return nil
	} else {
		return account.(*models.Account)
	}
}

func (m Account) Create(account models.Account) {
	m.Called(account)
}
