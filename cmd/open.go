package cmd

import (
	"errors"
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/session"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open an existing account or budget",
	Long:  `Open an existing account or budget`,
	RunE:  OpenCmd,
}

func OpenCmd(cmd *cobra.Command, args []string) error {
	// Init a session for the user
	session.InitUserSession()
	// Manually inject entities
	return OpenCmdDo(args, dao.UserDaoImpl{}, dao.AccountDaoImpl{},
		session.GetSession().GetUser())
}

func OpenCmdDo(args []string, userDao dao.UserDao, accountDao dao.AccountDao,
	user *entities.User) error {

	if len(args) == 0 {
		return errors.New("The name of the account to open should be provided.")
	}

	account, err := accountDao.FindByName(user.ID, args[0])

	if err != nil {
		return err
	}

	user.CliParams = entities.CliParams{
		CurrentAccountId: account.ID,
	}

	return userDao.Update(*user)
}

func init() {
	RootCmd.AddCommand(openCmd)
}
