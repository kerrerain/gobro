package expense

import (
	"bufio"
	"github.com/magleff/gobro/database"
	"log"
	"os"
	"strings"
	"time"
)

type ExpenseController struct {
	Datastore *ExpenseDataStore
}

func NewController(DB *database.Database) *ExpenseController {
	instance := new(ExpenseController)
	instance.Datastore = NewDatastore(DB)
	return instance
}

func (self ExpenseController) ImportFromFile(file *os.File) {
	expenses := extractFromFile(file)
	self.Datastore.ImportExpensesIntoDB(expenses)
}

func parseTime(input string) time.Time {
	time, err := time.Parse("02/01/2006", input)
	if err != nil {
		log.Fatal(err)
	}
	return time
}

func processLine(line string) Expense {
	fields := strings.Split(line, ";")
	return Expense{parseTime(fields[0]), fields[1], parseAmount(fields[2])}
}

func extractFromFile(file *os.File) []Expense {
	scanner := bufio.NewScanner(file)
	countLine := 0
	var expenses []Expense

	for scanner.Scan() {
		if countLine > 7 {
			expenses = append(expenses, processLine(scanner.Text()))
		}
		countLine++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return expenses
}
