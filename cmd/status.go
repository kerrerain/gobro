package cmd

import (
	"fmt"
	"github.com/magleff/gobro/controllers/budget"
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/dto"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/session"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gives the status of the current budget",
	Long:  `Gives the status of the current budget`,
	RunE:  StatusCmd,
}

func StatusCmd(cmd *cobra.Command, args []string) error {
	// Init a session for the user
	session.InitUserSession()
	// Manually inject entities
	return StatusCmdDo(args, dao.AccountDaoImpl{}, budget.BudgetControllerImpl{},
		session.GetSession().GetUser())
}

func StatusCmdDo(args []string, accountDao dao.AccountDao,
	budgetController budget.BudgetController, user *entities.User) error {

	account, err := accountDao.FindById(user.CliParams.CurrentAccountId)

	if err != nil {
		return err
	}

	fmt.Println("On account", account.Name)

	budgetInfo, errInfo := budgetController.ComputeInformation(user.CliParams.CurrentAccountId)

	if errInfo != nil {
		return errInfo
	}

	displayBudgetInfos(budgetInfo)

	return nil
}

func displayBudgetInfos(budgetInfo *dto.BudgetInformation) {
	fmt.Println("Created on", budgetInfo.StartDate)
	fmt.Println("Initial balance", budgetInfo.InitialBalance)
	fmt.Println("Total earnings", budgetInfo.TotalEarnings)
	fmt.Println("Total expenses", budgetInfo.TotalExpenses)
	fmt.Println("Total unchecked expenses", budgetInfo.TotalUncheckedExpenses)
	fmt.Println("Balance", budgetInfo.CurrentBalance.String(), "("+budgetInfo.Difference.String()+")")
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
