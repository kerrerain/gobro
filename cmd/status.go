package cmd

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/budget"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gives the status of the current budget",
	Long:  `Gives the status of the current budget`,
	Run: func(cmd *cobra.Command, args []string) {
		DB := database.NewDatabase()
		controller := budget.NewController(DB)
		controller.CurrentBudget()
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
