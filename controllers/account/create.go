package account

import (
	"errors"
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/session"
)

func (c AccountControllerImpl) Create(name string) error {
	// Manually inject entities
	return CreateDo(dao.AccountDaoImpl{}, session.GetCurrentUser(), name)
}

func CreateDo(accountDao dao.AccountDao, user *entities.User, name string) error {
	if len(name) == 0 {
		return errors.New("A name should be provided for the new account.")
	}

	_, err := accountDao.FindByName(name)

	// No error means that the entity was actually found
	if err == nil {
		return errors.New("An account already exists with the same name.")
	} else {
		accountDao.Create(*user, entities.Account{Name: name})
	}

	return nil
}
