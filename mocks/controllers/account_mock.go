package controllers

import (
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/mock"
)

type Account struct {
	mock.Mock
}

func (m Account) List() []models.Account {
	args := m.Called()
	if accounts := args.Get(0); accounts == nil {
		return nil
	} else {
		return accounts.([]models.Account)
	}
}

func (m Account) Create(name string) error {
	args := m.Called()
	if err := args.Error(0); err == nil {
		return nil
	} else {
		return err
	}
}
