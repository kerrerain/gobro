package controllers

import (
	"errors"
	"github.com/magleff/gobro/models"
)

type AccountController interface {
	List(models.AccountEntity) []models.Account
	Open(models.AccountEntity, string) error
}

type Account struct{}

func (c Account) List(entity models.AccountEntity) []models.Account {
	return entity.GetAll()
}

func (c Account) Open(entity models.AccountEntity, name string) error {
	if len(name) == 0 {
		return errors.New("A name should be provided for the new account.")
	}

	if existingAccount := entity.FindByName(name); existingAccount == nil {
		entity.Create(models.Account{Name: name})
	} else {
		return errors.New("An account already exists with this name.")
	}

	return nil
}
