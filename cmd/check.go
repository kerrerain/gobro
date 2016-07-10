package cmd

import (
	"fmt"
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/budget"
	"github.com/magleff/gobro/features/expense"
	"github.com/spf13/cobra"
)

// As seen in https://gist.github.com/albrow/5882501s

// askForConfirmation uses Scanln to parse user input. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user. Typically, you should use fmt to print out a question
// before calling askForConfirmation. E.g. fmt.Println("WARNING: Are you sure? (yes/no)")
func askForConfirmation() bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		return false
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	return containsString(okayResponses, response)
}

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// containsString returns true if slice contains element
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

func checkExpense(entry *expense.Expense) {
	fmt.Println("-> ", entry.Amount, entry.Description, entry.Date, "? (y/...)")
	if askForConfirmation() {
		entry.Checked = true
		fmt.Println("checked")
	} else {
		fmt.Println("not checked")
	}
}

var checkCommand = &cobra.Command{
	Use:   "check",
	Short: "Check the expenses",
	Long:  `Check the expenses`,
	Run: func(cmd *cobra.Command, args []string) {
		DB := database.NewDatabase()
		budgetController := budget.NewController(DB)
		currentBudget := budgetController.CurrentBudget()
		expenses := currentBudget.Expenses

		for index, entry := range expenses {
			if !entry.Checked {
				checkExpense(&expenses[index])
				currentBudget.Expenses = expenses
				budgetController.SaveBudget(currentBudget)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(checkCommand)
}
