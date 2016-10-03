package account

import (
	"github.com/magleff/gobro/entities"
)

type AccountController interface {
	List() []entities.Account
	Create(string) error
}

type AccountControllerImpl struct{}
