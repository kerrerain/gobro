package cmd

import (
	"github.com/magleff/gobro/features/budget"
	"github.com/spf13/cobra"
)

type GobroInitCommand struct {
	Command          *cobra.Command
	FlagPristine     bool
	BudgetController budget.BudgetController
}

func (self *GobroInitCommand) Init() {
	self.Command = &cobra.Command{
		Use:   "init",
		Short: "Init a new sheet for the budget",
		Long:  `Init a new sheet for the budget`,
		RunE:  self.Run,
	}
	self.Command.Flags().BoolVarP(&self.FlagPristine, "pristine", "p", false, "If this flag is set, "+
		"creates a budget without fixed expenses")
	self.BudgetController = budget.NewBudgetController()
}

func (self *GobroInitCommand) Run(cmd *cobra.Command, args []string) error {
	var err error
	balance := "0"

	if len(args) > 0 {
		balance = args[0]
	}

	if !self.FlagPristine {
		err = self.BudgetController.CreateBudgetWithFixedExpenses(balance)
	} else {
		err = self.BudgetController.CreatePristineBudget(balance)
	}

	return err
}

func init() {
	command := GobroInitCommand{}
	command.Init()
	RootCmd.AddCommand(command.Command)
}
