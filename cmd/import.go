package cmd

import (
	"fmt"
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/expense"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var filePath string

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Imports expenses from a .csv file",
	Long:  `Imports expenses from a .csv file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Import expenses from", filePath)

		DB := database.NewDatabase()
		file := openFile(filePath)
		controller := expense.NewController(DB)
		controller.ImportFromFile(file)

		defer file.Close()
	},
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
	RootCmd.AddCommand(importCmd)
}
