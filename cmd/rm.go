package cmd

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/expensefixed"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove something",
	Long:  `Remove something`,
	Run: func(cmd *cobra.Command, args []string) {
		session := database.CreateSession()
		controller := expensefixed.Controller(session)
		controller.RemoveExpenseFixed(args[0])

		defer session.Close()
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
