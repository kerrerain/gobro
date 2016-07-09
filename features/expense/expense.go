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
	Checked     bool
}

func NewExpense(amount string, description string) *Expense {
	instance := new(Expense)
	instance.Date = time.Now()
	instance.Description = description
	if strings.Contains(amount, "+") {
		instance.Amount = parseAmount(amount)
	} else {
		instance.Amount = parseAmount(amount) * -1
	}
	instance.Checked = false
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
