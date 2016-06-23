package expense

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	Date        time.Time
	Description string
	Amount      float32
}

func NewExpense(amount string, description string) *Expense {
	instance := new(Expense)
	instance.Date = time.Now()
	instance.Description = description
	instance.Amount = parseAmount(amount)
	instance.Amount = -1 * instance.Amount
	return instance
}

func NewResource(amount string, description string) *Expense {
	instance := new(Expense)
	instance.Date = time.Now()
	instance.Description = description
	instance.Amount = parseAmount(amount)
	return instance
}

func parseAmount(amount string) float32 {
	amount = strings.Replace(amount, ",", ".", 1)
	amountFloat, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		log.Fatal(err)
	}
	return float32(amountFloat)
}
