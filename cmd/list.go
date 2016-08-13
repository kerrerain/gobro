package cmd

import (
	"fmt"
	"github.com/magleff/gobro/features/account"
	"github.com/magleff/gobro/features/budget"
	"github.com/spf13/cobra"
)

type GobroListCommand struct {
	Command           *cobra.Command
	AccountController account.AccountController
	BudgetController  budget.BudgetController
}

func (self *GobroListCommand) Init() {
	self.Command = &cobra.Command{
		Use:   "list [account] or list [budget] or list",
		Short: "Lists the elements",
		Long: `Lists the accounts, the budgets created for the current account,
			or the expenses for the current budget if you simply run "list".`,
		RunE: self.Run,
	}
	self.AccountController = account.NewAccountController()
	self.BudgetController = budget.NewBudgetController()
}

func (self *GobroListCommand) Run(cmd *cobra.Command, args []string) error {
	var err error

	if len(args) == 0 {
		currentBudget := self.BudgetController.CurrentBudget()
		displayBudgetExpenses(*currentBudget)
	} else if typeOfElement := args[0]; typeOfElement == "account" {
		accounts := self.AccountController.List()
		displayAccounts(accounts)
	}

	return err
}

func printChecked(checked bool) string {
	return printBoolean(checked, "X")
}

func printActive(active bool) string {
	return printBoolean(active, "*")
}

func printBoolean(boolean bool, char string) string {
	str := ""
	if boolean {
		str = char
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

func displayAccounts(accounts []account.Account) {
	for _, entry := range accounts {
		fmt.Printf("%-2s %-30s\n", printActive(entry.Active), entry.Name)
	}
}

func init() {
	command := GobroListCommand{}
	command.Init()
	RootCmd.AddCommand(command.Command)
}
