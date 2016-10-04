package budget

import (
	"errors"
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/entities"
	"gopkg.in/mgo.v2/bson"
)

func (c BudgetControllerImpl) Create(userId bson.ObjectId, accountId bson.ObjectId) error {
	return CreateDo(dao.BudgetDaoImpl{},
		dao.AccountDaoImpl{}, userId, accountId)
}

func CreateDo(budgetDao dao.BudgetDao,
	accountDao dao.AccountDao, userId bson.ObjectId, accountId bson.ObjectId) error {

	_, findAccountError := accountDao.FindById(accountId)

	if findAccountError != nil {
		return errors.New("The account does not exist.")
	}

	_, activeBudgetNotFoundError := budgetDao.FindActiveBudget(accountId)

	// No error means that there's already an active budget,
	// so it is not possible to create a new one yet.
	if activeBudgetNotFoundError == nil {
		return errors.New("There is already an opened budget." +
			"Please close the current budget before creating a new one.")
	}

	return budgetDao.Create(entities.Budget{AccountId: accountId, UserId: userId})
}
