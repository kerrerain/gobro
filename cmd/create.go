package cmd

import (
	"errors"
	"github.com/magleff/gobro/controllers/account"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates something",
	Long:  `Creates something`,
	RunE:  CreateCmd,
}

func CreateCmd(cmd *cobra.Command, args []string) error {
	// Manually inject entities
	return CreateCmdDo(args, account.AccountControllerImpl{})
}

func CreateCmdDo(args []string, accountController account.AccountController) error {
	if len(args) == 0 {
		return errors.New("A name should be provided for the account.")
	}
	return accountController.Create(args[0])
}

func init() {
	RootCmd.AddCommand(createCmd)
}
