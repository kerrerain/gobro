package user

import (
	"errors"
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/session"
)

func (c UserControllerImpl) OpenAccount(userName string, accountName string) error {
	// Manually inject entities
	return OpenAccountDo(dao.UserDaoImpl{}, dao.AccountDaoImpl{},
		session.GetCurrentUser(), accountName)
}

func OpenAccountDo(userDao dao.UserDao, accountDao dao.AccountDao, user *entities.User,
	accountName string) error {

	if len(accountName) == 0 {
		return errors.New("The name of the account to open should be provided.")
	}

	account, err := accountDao.FindByName(accountName)

	if err != nil {
		return errors.New("This account doesn't exist. " + err.Error())
	}

	user.CurrentAccountId = account.ID
	if databaseErr := userDao.Update(*user); databaseErr != nil {
		return err
	}

	return nil
}
