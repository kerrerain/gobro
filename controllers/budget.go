package controllers

import (
	"github.com/magleff/gobro/dto"
	"github.com/magleff/gobro/models"
	"github.com/magleff/gobro/utils"
)

type BudgetController interface {
	ComputeInformation(models.BudgetEntity) *dto.BudgetInformation
}

type Budget struct{}

func (c Budget) ComputeInformation(entity models.BudgetEntity) *dto.BudgetInformation {
	information := new(dto.BudgetInformation)
	budget := entity.GetCurrent()

	information.StartDate = budget.StartDate
	information.InitialBalance = budget.InitialBalance
	information.TotalEarnings = utils.ComputeTotalEarnings(budget.Expenses)
	information.TotalExpenses = utils.ComputeTotalExpenses(budget.Expenses)
	information.TotalUncheckedExpenses = utils.ComputeTotalUncheckedExpenses(budget.Expenses)
	information.Difference = information.TotalEarnings.Add(information.TotalExpenses)
	information.CurrentBalance = information.InitialBalance.Add(information.Difference)

	return information
}
