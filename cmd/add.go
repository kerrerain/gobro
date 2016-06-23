package cmd

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/budget"
	"github.com/magleff/gobro/features/expensefixed"
	"github.com/spf13/cobra"
)

func parseArguments(args []string) (string, string) {
	var amount string
	var description string

	amount = args[0]

	if len(args) > 1 {
		description = args[1]
	}

	return amount, description
}

var typeOfExpense string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add something",
	Long:  `Add something`,
	Run: func(cmd *cobra.Command, args []string) {
		amount, description := parseArguments(args)
		DB := database.NewDatabase()

		if typeOfExpense == "fixed" {
			controller := expensefixed.NewController(DB)
			controller.CreateExpenseFixed(amount, description)
		} else {
			budgetController := budget.NewController(DB)
			budgetController.AddExpenseToCurrentBudget(amount, description)
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&typeOfExpense, "type", "t", "", "Type of the expense")
	RootCmd.AddCommand(addCmd)
}
