package account

import (
	"github.com/magleff/gobro/models"
)

type Controller interface {
	List(models.AccountEntity) []models.Account
	Create(models.AccountEntity, string) error
}

type Impl struct{}
