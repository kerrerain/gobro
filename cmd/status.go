package cmd

import (
	"fmt"
	"github.com/magleff/gobro/models"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gives the status of the current budget",
	Long:  `Gives the status of the current budget`,
	Run:   RunStatusCmd,
}

func RunStatusCmd(cmd *cobra.Command, args []string) {
	StatusCmd(args, models.Budget{})
}

func StatusCmd(args []string, budgetEntity models.BudgetEntity) {
	budget := budgetEntity.GetCurrent()
	fmt.Println(budget.InitialBalance)
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
