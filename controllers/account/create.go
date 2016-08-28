package account

import (
	"errors"
	"github.com/magleff/gobro/models"
	"github.com/magleff/gobro/session"
)

func (c Impl) Create(name string) error {
	// Manually inject entities
	return CreateDo(models.Account{}, session.GetCurrentUser(), name)
}

func CreateDo(entity models.AccountEntity, user *models.User, name string) error {
	if len(name) == 0 {
		return errors.New("A name should be provided for the new account.")
	}

	_, err := entity.FindByName(name)

	// No error means that the entity was actually found
	if err == nil {
		return errors.New("An account already exists with the same name.")
	} else {
		entity.Create(*user, models.Account{Name: name})
	}

	return nil
}
