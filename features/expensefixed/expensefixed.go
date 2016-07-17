package expensefixed

import (
	amountUtils "github.com/magleff/gobro/utils/amount"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strings"
	"time"
)

type ExpenseFixed struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Date        time.Time
	Description string
	Amount      float32
}

func parseAmount(amount string) float32 {
	parsedAmount, err := amountUtils.ParseString(amount)
	if !strings.Contains(amount, "+") {
		parsedAmount = parsedAmount * -1
	}
	if err != nil {
		log.Fatal(err)
	}
	return parsedAmount
}

func NewExpenseFixed(amount string, description string) *ExpenseFixed {
	instance := new(ExpenseFixed)
	instance.Date = time.Now()
	instance.Description = description
	instance.Amount = parseAmount(amount)
	return instance
}
