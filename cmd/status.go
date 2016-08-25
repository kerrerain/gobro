package cmd

import (
	"errors"
	"github.com/magleff/gobro/models"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gives the status of the current budget",
	Long:  `Gives the status of the current budget`,
	RunE:  RunStatusCmd,
}

func RunStatusCmd(cmd *cobra.Command, args []string) error {
	return StatusCmd(args, models.Budget{})
}

func StatusCmd(args []string, budgetEntity models.BudgetEntity) error {
	var err error

	budget := budgetEntity.GetCurrent()

	if budget != nil {
		budget.Display()
	} else {
		err = errors.New("There is not any active budget. " +
			"use 'open budget' to open a new budget.")
	}

	return err
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
