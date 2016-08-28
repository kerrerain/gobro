package user

import (
	"errors"
	"github.com/magleff/gobro/models"
	"github.com/magleff/gobro/session"
)

func (c Impl) OpenAccount(userName string, accountName string) error {
	// Manually inject entities
	return OpenAccountDo(models.User{}, models.Account{}, session.GetCurrentUser(), accountName)
}

func OpenAccountDo(userEntity models.UserEntity, accountEntity models.AccountEntity, user *models.User,
	accountName string) error {

	if len(accountName) == 0 {
		return errors.New("The name of the account to open should be provided.")
	}

	account, err := accountEntity.FindByName(accountName)

	if err != nil {
		return errors.New("This account doesn't exist. " + err.Error())
	}

	userEntity.UpdateAccount(*user, *account)

	return nil
}
