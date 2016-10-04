package account

import (
	"errors"
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/entities"
	"gopkg.in/mgo.v2/bson"
)

func (c AccountControllerImpl) Create(userId bson.ObjectId, name string) error {
	// Manually inject entities
	return CreateDo(dao.AccountDaoImpl{}, userId, name)
}

func CreateDo(accountDao dao.AccountDao, userId bson.ObjectId, name string) error {
	if len(name) == 0 {
		return errors.New("A name should be provided for the new account.")
	}

	_, err := accountDao.FindByName(userId, name)

	// No error means that the entity was actually found
	if err == nil {
		return errors.New("An account already exists with the same name.")
	} else {
		accountDao.Create(entities.Account{Name: name, UserId: userId})
	}

	return nil
}
