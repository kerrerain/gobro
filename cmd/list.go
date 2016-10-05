package cmd

import (
	"fmt"
	"github.com/magleff/gobro/controllers/account"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/session"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List something",
	Long:  `List something`,
	RunE:  ListCmd,
}

func ListCmd(cmd *cobra.Command, args []string) error {
	// Init a session for the user
	session.InitUserSession()
	// Manually inject entities
	return ListCmdDo(args, account.AccountControllerImpl{}, session.GetSession().GetUser())
}

func ListCmdDo(args []string, accountController account.AccountController,
	user *entities.User) error {

	accounts, err := accountController.List(user.ID)

	if err != nil {
		return err
	}

	for _, entry := range accounts {
		fmt.Println(entry.Name)
	}

	return nil
}

func init() {
	RootCmd.AddCommand(listCmd)
}
