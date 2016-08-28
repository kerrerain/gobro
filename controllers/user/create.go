package user

import (
	"errors"
	"github.com/magleff/gobro/models"
)

func (c Impl) Create(userName string) error {
	// Manually inject entities
	return CreateDo(models.User{}, userName)
}

func CreateDo(entity models.UserEntity, userName string) error {
	_, err := entity.FindByName(userName)

	// No error means that the entity was actually found
	if err == nil {
		return errors.New("This user already exists.")
	} else {
		entity.Create(models.User{Name: userName})
	}

	return nil
}
