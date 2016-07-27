package cmd

import (
	"github.com/magleff/gobro/features/budget"
	"github.com/spf13/cobra"
)

var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Close the budget",
	Long:  `Close the budget`,
	Run: func(cmd *cobra.Command, args []string) {
		controller := budget.NewBudgetController()
		controller.CloseCurrentBudget()
	},
}

func init() {
	RootCmd.AddCommand(closeCmd)
}
