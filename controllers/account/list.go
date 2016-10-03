package account

import (
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/entities"
)

func (c AccountControllerImpl) List() []entities.Account {
	// Manually inject entities
	return ListDo(dao.AccountDaoImpl{})
}

func ListDo(accountDao dao.AccountDao) []entities.Account {
	return accountDao.GetAll()
}
