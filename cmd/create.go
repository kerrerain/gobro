package cmd

import (
	"errors"
	"github.com/magleff/gobro/common"
	"github.com/magleff/gobro/controllers/account"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/session"
	"github.com/spf13/cobra"
	"gopkg.in/mgo.v2/bson"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates something",
	Long:  `Creates something`,
	RunE:  CreateCmd,
}

func CreateCmd(cmd *cobra.Command, args []string) error {
	// Init a session for the user
	session.InitUserSession()
	// Manually inject entities
	return CreateCmdDo(args, account.AccountControllerImpl{}, session.GetSession().GetUser())
}

func CreateCmdDo(args []string, accountController account.AccountController,
	user *entities.User) error {

	if len(args) == 0 {
		return errors.New("A type should be provided for the object to create (account or budget).")
	}

	if args[0] == common.TYPE_ACCOUNT {
		return createAccount(args, accountController, user.ID)
	}

	return nil
}

func createAccount(args []string, accountController account.AccountController,
	userId bson.ObjectId) error {
	if len(args) == 1 {
		return errors.New("A name should be provided for the account.")
	} else {
		return accountController.Create(userId, args[1])
	}
}

func init() {
	RootCmd.AddCommand(createCmd)
}
