package cmd

import (
	"errors"
	"github.com/magleff/gobro/features/budget"
	"github.com/spf13/cobra"
)

type GobroInitCommand struct {
	Command          *cobra.Command
	FlagFixed        bool
	BudgetController budget.BudgetController
}

func (self *GobroInitCommand) Init() {
	self.Command = &cobra.Command{
		Use:   "init [initial balance]",
		Short: "Init a new sheet for the budget",
		Long: `Init a new sheet for the budget. You must provide the initial balance of the budget
			(for example: 100.38)`,
		RunE: self.Run,
	}
	self.Command.Flags().BoolVarP(&self.FlagFixed, "fixed", "f", false, "If this flag is set, "+
		"creates a budget with fixed expenses")
	self.BudgetController = budget.NewBudgetController()
}

func (self *GobroInitCommand) Run(cmd *cobra.Command, args []string) error {
	var err error
	var balance string

	if len(args) > 0 {
		balance = args[0]
	} else {
		return errors.New("You should provide the initial balance of the budget " +
			"(for example: 100.38).")
	}

	if self.FlagFixed {
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
