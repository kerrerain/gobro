package budget

import (
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/dto"
	"github.com/magleff/gobro/utils"
	"gopkg.in/mgo.v2/bson"
)

func (c BudgetControllerImpl) ComputeInformation(accountId bson.ObjectId) (*dto.BudgetInformation, error) {
	return ComputeInformationDo(dao.BudgetDaoImpl{}, accountId)
}

func ComputeInformationDo(budgetDao dao.BudgetDao,
	accountId bson.ObjectId) (*dto.BudgetInformation, error) {

	information := new(dto.BudgetInformation)
	budget, err := budgetDao.FindActiveBudget(accountId)

	information.StartDate = budget.StartDate
	information.InitialBalance = budget.InitialBalance
	information.TotalEarnings = utils.ComputeTotalEarnings(budget.Expenses)
	information.TotalExpenses = utils.ComputeTotalExpenses(budget.Expenses)
	information.TotalUncheckedExpenses = utils.ComputeTotalUncheckedExpenses(budget.Expenses)
	information.Difference = information.TotalEarnings.Add(information.TotalExpenses)
	information.CurrentBalance = information.InitialBalance.Add(information.Difference)

	return information, err
}
