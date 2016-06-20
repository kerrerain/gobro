package expensefixed

import (
	"gopkg.in/mgo.v2"
	"log"
	"strconv"
	"strings"
)

type ExpenseFixedController struct {
	session *mgo.Session
}

func Controller(session *mgo.Session) *ExpenseFixedController {
	return &ExpenseFixedController{session}
}

func (ec ExpenseFixedController) CreateExpenseFixed(amount string, description string) {
	DataStore(ec.session).CreateExpenseFixed(parseAmount(amount), description)
}

func (ec ExpenseFixedController) ListExpensesFixed() []ExpenseFixed {
	return DataStore(ec.session).ListExpensesFixed()
}

func (ec ExpenseFixedController) RemoveExpenseFixed(index string) {
	DataStore(ec.session).RemoveExpenseFixed(parseIndex(index))
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

func parseIndex(index string) int32 {
	parsedIndex, err := strconv.ParseInt(index, 0, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int32(parsedIndex)
}
