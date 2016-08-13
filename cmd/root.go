// This package contains the cobra commands of the CLI application.
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var filePath string
var typeOfExpense string
var pristine bool
var mailFlag bool

var RootCmd = &cobra.Command{
	Use:   "gobro",
	Short: "Simple budget management",
	Long: `It stands for "Go" and "Slowbro", which is a lovely pokemon -- but quite lazy.
Since I am lazy too when it comes to making a budget, Gobro, I choose you!`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO get the current status of the database (current budget, etc).
		fmt.Println("Current status: OK")
	},
}

func init() {}
