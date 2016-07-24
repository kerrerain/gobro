package expense

import (
	amountUtils "github.com/magleff/gobro/utils/amount"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strings"
	"time"
)

type Expense struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Date        time.Time
	Description string
	Amount      float32
	Checked     bool
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

func NewExpense(amount string, description string) *Expense {
	instance := new(Expense)
	instance.Date = time.Now()
	instance.Description = description
	instance.Amount = parseAmount(amount)
	instance.Checked = false
	return instance
}
