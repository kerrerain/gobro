package main

import (
	"flag"
	"github.com/magleff/gobro/expense"
	"log"
	"os"
)

func parseFileName() string {
	fileFlag := flag.String("f", "", "The path to the file")
	flag.Parse()
	return *fileFlag
}

func main() {
	file, err := os.Open(parseFileName())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	expense.ImportFromFile(file)
}
