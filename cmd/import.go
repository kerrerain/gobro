package cmd

import (
	"fmt"
	"github.com/magleff/gobro/features/budget"
	"github.com/magleff/gobro/features/expense"
	"github.com/magleff/gobro/features/mail"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Imports expenses from a .csv file",
	Long:  `Imports expenses from a .csv file`,
	Run: func(cmd *cobra.Command, args []string) {
		if mailFlag {
			importFromMail()
		} else {
			fmt.Println("Import expenses from", filePath)
			importFromFile(filePath)
		}
	},
}

func importFromMail() {
	controller := mail.NewMailController()
	budgetController := budget.NewBudgetController()
	expenses := controller.ImportFromMail()
	budgetController.AddRawExpensesToCurrentBudget(expenses)
}

func importFromFile(filePath string) {
	file := openFile(filePath)
	controller := expense.NewExpenseController()
	controller.ImportFromFile(file)
	defer file.Close()
}

func openFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func init() {
	importCmd.Flags().StringVarP(&filePath, "filepath", "f", "", "Path to the file to import")
	importCmd.Flags().BoolVarP(&mailFlag, "mail", "m", false, "Import the last expenses sent by mail")
	RootCmd.AddCommand(importCmd)
}
