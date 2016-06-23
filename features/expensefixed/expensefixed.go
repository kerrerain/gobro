package expensefixed

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"strconv"
	"strings"
	"time"
)

type ExpenseFixed struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Date        time.Time
	Description string
	Amount      float32
}

func NewExpenseFixed(amount string, description string) *ExpenseFixed {
	instance := new(ExpenseFixed)
	instance.Date = time.Now()
	instance.Description = description
	instance.Amount = parseAmount(amount)
	instance.Amount = -1 * instance.Amount
	return instance
}

// FIXME duplicate code
func parseAmount(amount string) float32 {
	amount = strings.Replace(amount, ",", ".", 1)
	amountFloat, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		log.Fatal(err)
	}
	return float32(amountFloat)
}
