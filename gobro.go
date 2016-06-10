package main

import (
	"flag"
	"fmt"
	"github.com/magleff/gobro/expense"
	"log"
	"os"
)

func openFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func main() {
	commandFlag := flag.String("c", "", "The command to run")
	fileFlag := flag.String("f", "", "The path to the file")
	flag.Parse()

	switch *commandFlag {
	case "add-fixed":
		fmt.Println("Add a fixed expense")
	case "import":
		fmt.Println("Import expenses from file")
		file := openFile(*fileFlag)
		expense.ImportFromFile(file)
		defer file.Close()
	default:
		fmt.Println("Unrecognized command", commandFlag)
	}

	fmt.Println(*fileFlag)
}
