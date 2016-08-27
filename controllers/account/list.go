package account

import (
	"github.com/magleff/gobro/models"
)

func (c Impl) List(entity models.AccountEntity) []models.Account {
	return entity.GetAll()
}
