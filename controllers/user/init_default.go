package user

import (
	"github.com/magleff/gobro/common"
)

func (c UserControllerImpl) InitDefault() error {
	// Manually inject entities
	return InitDefaultDo(UserControllerImpl{})
}

func InitDefaultDo(userController UserController) error {
	return userController.Create(common.DEFAULT_USER_NAME)
}
