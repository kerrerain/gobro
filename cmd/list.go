package cmd

import (
	"fmt"
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/expensefixed"
	"github.com/spf13/cobra"
)

func displayExpensesFixed(expensesFixed []expensefixed.ExpenseFixed) {
	for index, entry := range expensesFixed {
		fmt.Println(index, ")", entry.Amount, entry.Description, entry.Date)
	}
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List something",
	Long:  `List something`,
	Run: func(cmd *cobra.Command, args []string) {
		session := database.CreateSession()
		controller := expensefixed.Controller(session)
		expensesFixed := controller.ListExpensesFixed()
		displayExpensesFixed(expensesFixed)

		defer session.Close()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
