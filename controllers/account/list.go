package account

import (
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/entities"
	"gopkg.in/mgo.v2/bson"
)

func (c AccountControllerImpl) List(userId bson.ObjectId) ([]entities.Account, error) {
	// Manually inject entities
	return ListDo(dao.AccountDaoImpl{}, userId)
}

func ListDo(accountDao dao.AccountDao, userId bson.ObjectId) ([]entities.Account, error) {
	return accountDao.GetAll(userId)
}
