package expensefixed

import (
	"github.com/magleff/gobro/features/expense"
	"log"
	"strconv"
)

type ExpenseFixedController struct {
	ExpenseFixedDatastore *ExpenseFixedDatastore
}

func NewExpenseFixedController() *ExpenseFixedController {
	instance := new(ExpenseFixedController)
	instance.ExpenseFixedDatastore = new(ExpenseFixedDatastore)
	return instance
}

func (self ExpenseFixedController) CreateExpenseFixed(amount string, description string) {
	self.ExpenseFixedDatastore.CreateExpenseFixed(*expense.NewExpense(amount, description))
}

func (self ExpenseFixedController) ListExpensesFixed() []expense.Expense {
	return self.ExpenseFixedDatastore.ListExpensesFixed()
}

func (self ExpenseFixedController) RemoveExpenseFixed(index string) {
	self.ExpenseFixedDatastore.RemoveExpenseFixed(parseIndex(index))
}

func parseIndex(index string) int32 {
	parsedIndex, err := strconv.ParseInt(index, 0, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int32(parsedIndex)
}
