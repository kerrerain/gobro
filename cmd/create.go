package cmd

import (
	"errors"
	"github.com/magleff/gobro/common"
	"github.com/magleff/gobro/controllers/account"
	"github.com/magleff/gobro/controllers/budget"
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
	return CreateCmdDo(args, account.AccountControllerImpl{}, budget.BudgetControllerImpl{},
		session.GetSession().GetUser())
}

func CreateCmdDo(args []string, accountController account.AccountController,
	budgetController budget.BudgetController, user *entities.User) error {

	if len(args) == 0 {
		return errors.New("A type should be provided for the object to create (account or budget).")
	}

	if args[0] == common.TYPE_ACCOUNT {
		return createAccount(args, accountController, user.ID)
	} else if args[0] == common.TYPE_BUDGET {
		return createBudget(budgetController, user)
	}

	return errors.New("Unkown type of object to create.")
}

func createAccount(args []string, accountController account.AccountController,
	userId bson.ObjectId) error {
	if len(args) == 1 {
		return errors.New("A name should be provided for the account.")
	} else {
		return accountController.Create(userId, args[1])
	}
}

func createBudget(budgetController budget.BudgetController, user *entities.User) error {
	return budgetController.Create(user.ID, user.CliParams.CurrentAccountId)
}

func init() {
	RootCmd.AddCommand(createCmd)
}
