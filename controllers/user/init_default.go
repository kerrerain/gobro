package user

import (
	"github.com/magleff/gobro/common"
	"github.com/magleff/gobro/models"
)

func (c Impl) InitDefault() {
	// Manually inject entities
	InitDefaultDo(models.User{})
}

func InitDefaultDo(entity models.UserEntity) {
	CreateDo(entity, common.DEFAULT_USER_NAME)
}
