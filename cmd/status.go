package cmd

import (
	"fmt"
	"github.com/magleff/gobro/features/budget"
	"github.com/spf13/cobra"
)

type GobroStatusCommand struct {
	Command          *cobra.Command
	BudgetController budget.BudgetController
}

func (self *GobroStatusCommand) Init() {
	self.Command = &cobra.Command{
		Use:   "status",
		Short: "Gives the status of the current budget",
		Long:  `Gives the status of the current budget`,
		RunE:  self.Run,
	}
	self.BudgetController = budget.NewBudgetController()
}

func (self *GobroStatusCommand) Run(cmd *cobra.Command, args []string) error {
	budgetInfo, err := self.BudgetController.ComputeBudgetInfo()
	if err == nil {
		displayBudgetInfos(budgetInfo)
	} else {
		return err
	}
	return nil
}

func displayBudgetInfos(budgetInfo *budget.BudgetInfo) {
	fmt.Println("Created on", budgetInfo.StartDate)
	fmt.Println("Initial balance", budgetInfo.InitialBalance)
	fmt.Println("Total earnings", budgetInfo.TotalEarnings)
	fmt.Println("Total expenses", budgetInfo.TotalExpenses)
	fmt.Println("Total unchecked expenses", budgetInfo.TotalUncheckedExpenses)
	fmt.Println("Balance", budgetInfo.CurrentBalance.String(), "("+budgetInfo.Difference.String()+")")
}

func init() {
	command := GobroStatusCommand{}
	command.Init()
	RootCmd.AddCommand(command.Command)
}
