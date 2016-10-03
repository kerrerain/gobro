package cmd

import (
	"fmt"
	"github.com/magleff/gobro/controllers/account"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List something",
	Long:  `List something`,
	Run:   ListCmd,
}

func ListCmd(cmd *cobra.Command, args []string) {
	// Manually inject entities
	ListCmdDo(args, account.AccountControllerImpl{})
}

func ListCmdDo(args []string, accountController account.AccountController) {
	accounts := accountController.List()
	for _, entry := range accounts {
		fmt.Println(entry.Name)
	}
}

func init() {
	RootCmd.AddCommand(listCmd)
}
