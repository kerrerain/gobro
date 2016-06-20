package cmd

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/budget"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a new sheet for the budget",
	Long:  `Init a new sheet for the budget`,
	Run: func(cmd *cobra.Command, args []string) {
		session := database.CreateSession()
		controller := budget.Controller(session)
		controller.CreateBudget()

		defer session.Close()
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
