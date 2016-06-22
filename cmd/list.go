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
		DB := database.NewDatabase()
		controller := expensefixed.NewController(DB)
		expensesFixed := controller.ListExpensesFixed()
		displayExpensesFixed(expensesFixed)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
