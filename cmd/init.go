package cmd

import (
	"github.com/magleff/gobro/budget"
	"github.com/magleff/gobro/database"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a new sheet for the budget",
	Long:  `Init a new sheet for the budget`,
	Run: func(cmd *cobra.Command, args []string) {
		database := database.NewDatabase()
		controller := budget.NewController(database)
		controller.CreateBudget()
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
