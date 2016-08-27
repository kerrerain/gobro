package account

import (
	"errors"
	"github.com/magleff/gobro/models"
)

func (c Impl) Create(name string) error {
	// Manually inject entities
	return CreateDo(models.Account{}, name)
}

func CreateDo(entity models.AccountEntity, name string) error {
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
