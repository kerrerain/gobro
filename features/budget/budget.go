package budget

import (
	"github.com/magleff/gobro/features/expense"
	"github.com/magleff/gobro/features/expensefixed"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strconv"
	"strings"
	"time"
)

type Budget struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	StartDate      time.Time
	EndDate        time.Time
	Expenses       []expense.Expense
	InitialBalance float32
	Active         bool
}

func NewBudget(expensesFixed []expensefixed.ExpenseFixed, balance string) *Budget {
	instance := new(Budget)
	instance.Expenses = convertExpensesFixed(expensesFixed)
	instance.StartDate = time.Now()
	instance.Active = true
	instance.InitialBalance = parseAmount(balance)
	return instance
}

func convertExpensesFixed(expensesFixed []expensefixed.ExpenseFixed) []expense.Expense {
	var expenses []expense.Expense
	for _, entry := range expensesFixed {
		expenses = append(expenses, expense.Expense{time.Now(), entry.Description, entry.Amount, false})
	}
	return expenses
}

func parseAmount(amount string) float32 {
	amount = strings.Replace(amount, ",", ".", 1)
	amountFloat, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		log.Fatal(err)
	}
	return float32(amountFloat)
}
