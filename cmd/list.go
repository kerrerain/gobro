package cmd

import (
	"fmt"
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/budget"
	"github.com/magleff/gobro/features/expensefixed"
	"github.com/spf13/cobra"
)

func displayExpensesFixed(expensesFixed []expensefixed.ExpenseFixed) {
	for index, entry := range expensesFixed {
		fmt.Println(index, ")", entry.Amount, entry.Description, entry.Date)
	}
}

func displayBudgetExpenses(budget budget.Budget) {
	for index, entry := range budget.Expenses {
		fmt.Println(index, ")", entry.Amount, entry.Description, entry.Date)
	}
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List something",
	Long:  `List something`,
	Run: func(cmd *cobra.Command, args []string) {
		DB := database.NewDatabase()

		if typeOfExpense == "fixed" {
			controller := expensefixed.NewController(DB)
			expensesFixed := controller.ListExpensesFixed()
			displayExpensesFixed(expensesFixed)
		} else {
			budgetController := budget.NewController(DB)
			currentBudget := budgetController.CurrentBudget()
			displayBudgetExpenses(*currentBudget)
		}
	},
}

func init() {
	listCmd.Flags().StringVarP(&typeOfExpense, "type", "t", "", "Type of the expense")
	RootCmd.AddCommand(listCmd)
}
