package budget

import (
	"github.com/magleff/gobro/dto"
	"github.com/magleff/gobro/utils"
)

func (c Impl) ComputeInformation() (*dto.BudgetInformation, error) {
	return ComputeInformationDo(Impl{})
}

func ComputeInformationDo(controller Controller) (*dto.BudgetInformation, error) {
	information := new(dto.BudgetInformation)
	budget, err := controller.Current()

	information.StartDate = budget.StartDate
	information.InitialBalance = budget.InitialBalance
	information.TotalEarnings = utils.ComputeTotalEarnings(budget.Expenses)
	information.TotalExpenses = utils.ComputeTotalExpenses(budget.Expenses)
	information.TotalUncheckedExpenses = utils.ComputeTotalUncheckedExpenses(budget.Expenses)
	information.Difference = information.TotalEarnings.Add(information.TotalExpenses)
	information.CurrentBalance = information.InitialBalance.Add(information.Difference)

	return information, err
}
