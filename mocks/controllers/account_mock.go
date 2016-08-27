package controllers

import (
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/mock"
)

type Account struct {
	mock.Mock
}

func (m Account) List(entity models.AccountEntity) []models.Account {
	args := m.Called(entity)
	if accounts := args.Get(0); accounts == nil {
		return nil
	} else {
		return accounts.([]models.Account)
	}
}

func (m Account) Create(entity models.AccountEntity, name string) error {
	args := m.Called(entity)
	if err := args.Error(0); err == nil {
		return nil
	} else {
		return err
	}
}
