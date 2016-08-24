package cmd

import (
	"fmt"
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
	ListCmd(args, models.Account{})
}

func ListCmd(args []string, accountEntity models.AccountEntity) {
	accounts := accountEntity.GetAll()
	for _, entry := range accounts {
		fmt.Println(entry.Name)
	}
}

func init() {
	RootCmd.AddCommand(listCmd)
}
