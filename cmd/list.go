package cmd

import (
	"fmt"
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/budget"
	"github.com/magleff/gobro/features/expense"
	"github.com/magleff/gobro/features/expensefixed"
	"github.com/spf13/cobra"
)

func displayExpensesFixed(expensesFixed []expense.Expense) {
	for index, entry := range expensesFixed {
		fmt.Println(index, ")", entry.Amount, entry.Description, entry.Date)
	}
}

func printChecked(checked bool) string {
	str := ""
	if checked {
		str = "X"
	}
	return str
}

func displayBudgetExpenses(budget budget.Budget) {
	fmt.Printf("\n")
	fmt.Printf("%-2s|%-6s|%-16s|%-30s|%-16s\n", "C", "Index", "Amount", "Description", "Date")
	fmt.Printf("\n")
	for index, entry := range budget.Expenses {
		fmt.Printf("%-2s|%-6v|%-16v|%-30s|%-16s\n",
			printChecked(entry.Checked),
			index,
			entry.Amount,
			entry.Description,
			entry.Date.Format("2006-01-02"))
	}
	fmt.Printf("\n")
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
