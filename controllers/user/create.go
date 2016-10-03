package user

import (
	"errors"
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/entities"
)

func (c UserControllerImpl) Create(userName string) error {
	// Manually inject entities
	return CreateDo(dao.UserDaoImpl{}, userName)
}

func CreateDo(userDao dao.UserDao, userName string) error {
	_, err := userDao.FindByName(userName)

	// No error means that the entity was actually found
	if err == nil {
		return errors.New("This user already exists.")
	} else {
		userDao.Create(entities.User{Name: userName})
	}

	return nil
}
