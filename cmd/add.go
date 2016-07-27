package cmd

import (
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

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add something",
	Long:  `Add something`,
	Run: func(cmd *cobra.Command, args []string) {
		amount, description := parseArguments(args)

		if typeOfExpense == "fixed" {
			controller := expensefixed.NewExpenseFixedController()
			controller.CreateExpenseFixed(amount, description)
		} else {
			budgetController := budget.NewBudgetController()
			budgetController.AddExpenseToCurrentBudget(amount, description)
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&typeOfExpense, "type", "t", "", "Type of the expense")
	RootCmd.AddCommand(addCmd)
}
