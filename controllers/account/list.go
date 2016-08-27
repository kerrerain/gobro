package account

import (
	"github.com/magleff/gobro/models"
)

func (c Impl) List() []models.Account {
	// Manually inject entities
	return ListDo(models.Account{})
}

func ListDo(entity models.AccountEntity) []models.Account {
	return entity.GetAll()
}
