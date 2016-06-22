package cmd

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/expensefixed"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove something",
	Long:  `Remove something`,
	Run: func(cmd *cobra.Command, args []string) {
		DB := database.NewDatabase()
		controller := expensefixed.NewController(DB)
		controller.RemoveExpenseFixed(args[0])
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
