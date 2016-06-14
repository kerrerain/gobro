package expensefixed

import (
	"gopkg.in/mgo.v2"
	"strconv"
	"strings"
	"log"
)

type ExpenseFixedController struct {
	session *mgo.Session
}

func Controller(session *mgo.Session) *ExpenseFixedController {
	return &ExpenseFixedController{session}
}

func (ec ExpenseFixedController) CreateExpenseFixed(description string, amount string) {
	dataStore(ec.session).CreateExpenseFixed(description, parseAmount(amount))
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
