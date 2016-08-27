package account

import (
	"github.com/magleff/gobro/models"
)

type Controller interface {
	List() []models.Account
	Create(string) error
}

type Impl struct{}
