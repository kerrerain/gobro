package main

import (
	"fmt"
	"github.com/magleff/gobro/cmd"
	"github.com/magleff/gobro/database"
	"os"
)

func main() {
	database.InitDatabase()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
