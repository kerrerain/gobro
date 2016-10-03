package user

import (
	"github.com/magleff/gobro/common"
	"github.com/magleff/gobro/dao"
)

func (c UserControllerImpl) InitDefault() {
	// Manually inject entities
	InitDefaultDo(dao.UserDaoImpl{})
}

func InitDefaultDo(userDao dao.UserDao) {
	CreateDo(userDao, common.DEFAULT_USER_NAME)
}
