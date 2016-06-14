package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/expensefixed"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add something",
	Long:  `Add something`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add", args[0])
		session := database.CreateSession()
		controller := expensefixed.Controller(session)
		controller.CreateExpenseFixed(args[1], args[2])
		defer session.Close()
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
