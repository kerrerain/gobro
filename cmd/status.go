package cmd

import (
	"fmt"
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/budget"
	"github.com/spf13/cobra"
)

func computeEarnings(budget budget.Budget) float32 {
	earnings := float32(0.00)
	for _, entry := range budget.Expenses {
		if entry.Amount > 0 {
			earnings += entry.Amount
		}
	}
	return earnings
}

func computeExpenses(budget budget.Budget) float32 {
	expenses := float32(0.00)
	for _, entry := range budget.Expenses {
		if entry.Amount <= 0 {
			expenses += entry.Amount
		}
	}
	return expenses
}

func displayBudgetInfos(budget budget.Budget) {
	earnings := computeEarnings(budget)
	expenses := computeExpenses(budget)

	fmt.Println("Created on", budget.StartDate)
	fmt.Println("Total earnings", earnings)
	fmt.Println("Total expenses", expenses)
	fmt.Println("Balance", earnings+expenses)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gives the status of the current budget",
	Long:  `Gives the status of the current budget`,
	Run: func(cmd *cobra.Command, args []string) {
		DB := database.NewDatabase()
		controller := budget.NewController(DB)
		currentBudget := controller.CurrentBudget()
		displayBudgetInfos(*currentBudget)
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
