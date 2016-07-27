package expense

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

type ExpenseController struct {
	ExpenseDatastore *ExpenseDatastore
}

func NewExpenseController() *ExpenseController {
	instance := new(ExpenseController)
	instance.ExpenseDatastore = new(ExpenseDatastore)
	return instance
}

func (self *ExpenseController) ImportFromFile(file *os.File) {
	expenses := extractFromFile(file)
	self.ExpenseDatastore.ImportExpensesIntoDB(expenses)
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
	expense := NewExpense(fields[2], fields[1])
	expense.Date = parseTime(fields[0])
	return *expense
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
