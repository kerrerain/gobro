package cmd

import (
	"github.com/magleff/gobro/features/budget"
	"github.com/spf13/cobra"
)

type GobroCloseCommand struct {
	Command          *cobra.Command
	BudgetController budget.BudgetController
}

func (self *GobroCloseCommand) Init() {
	self.Command = &cobra.Command{
		Use:   "close",
		Short: "Closes the current budget",
		Long:  `Closes the current budget`,
		RunE:  self.Run,
	}
	self.BudgetController = budget.NewBudgetController()
}

func (self *GobroCloseCommand) Run(cmd *cobra.Command, args []string) error {
	self.BudgetController.CloseCurrentBudget()
	return nil
}

func init() {
	command := GobroCloseCommand{}
	command.Init()
	RootCmd.AddCommand(command.Command)
}
