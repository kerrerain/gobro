package user

import (
	"errors"
	"github.com/magleff/gobro/models"
)

func (c Impl) Open(userName string, accountName string) error {
	// Manually inject entities
	return OpenDo(models.User{}, models.Account{}, userName, accountName)
}

func OpenDo(userEntity models.UserEntity, accountEntity models.AccountEntity, userName string,
	accountName string) error {
	if len(userName) == 0 {
		return errors.New("The name of the user should be provided.")
	}

	if len(accountName) == 0 {
		return errors.New("The name of the account to open should be provided.")
	}

	if account := accountEntity.FindByName(accountName); account == nil {
		return errors.New("This account doesn't exist.")
	}

	user := userEntity.FindByName(userName)

	if user == nil {
		return errors.New("This user doesn't exist.")
	}

	user.CurrentAccountName = accountName
	userEntity.Update(*user)

	return nil
}
