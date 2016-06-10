package expense

import (
	"bufio"
	"github.com/magleff/gobro/dbsession"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	Date        time.Time
	Description string
	Amount      float32
}

func parseTime(input string) time.Time {
	time, err := time.Parse("02/01/2006", input)
	if err != nil {
		log.Fatal(err)
	}
	return time
}

func parseAmount(amount string) float32 {
	amount = strings.Replace(amount, ",", ".", 1)
	amountFloat, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		log.Fatal(err)
	}
	return float32(amountFloat)
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

func importExpensesIntoDB(mongoSession *mgo.Session, entries []Expense) {
	expenses := mongoSession.DB("").C("expenses")
	for _, Expense := range entries {
		expenses.Insert(Expense)
	}
}

func ImportFromFile(file *os.File) {
	expenses := extractFromFile(file)
	mongoSession := dbsession.Create()

	importExpensesIntoDB(mongoSession, expenses)
}
