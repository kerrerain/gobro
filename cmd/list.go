package cmd

import (
	"fmt"
	"github.com/magleff/gobro/controllers/account"
	"github.com/magleff/gobro/models"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List something",
	Long:  `List something`,
	Run:   RunListCmd,
}

func RunListCmd(cmd *cobra.Command, args []string) {
	ListCmd(args, account.Impl{}, models.Account{})
}

func ListCmd(args []string, accountController account.Controller, accountEntity models.AccountEntity) {
	accounts := accountController.List(accountEntity)
	for _, entry := range accounts {
		fmt.Println(entry.Name)
	}
}

func init() {
	RootCmd.AddCommand(listCmd)
}
