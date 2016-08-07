package budget

import (
	"github.com/magleff/gobro/features/expense"
	"github.com/shopspring/decimal"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Budget struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	StartDate      time.Time
	EndDate        time.Time
	Expenses       []expense.Expense
	InitialBalance decimal.Decimal
	Active         bool
}

func (self *Budget) GetBSON() (interface{}, error) {
	initialBalanceFloat, _ := self.InitialBalance.Float64()

	return &struct {
		InitialBalance float64           `json:"initialbalance" bson:"initialbalance"`
		StartDate      time.Time         `json:"startdate" bson:"startdate"`
		EndDate        time.Time         `json:"enddate" bson:"enddate"`
		Expenses       []expense.Expense `json:"expenses" bson:"expenses"`
		Active         bool              `json:"active" bson:"active"`
	}{
		InitialBalance: initialBalanceFloat,
		StartDate:      self.StartDate,
		EndDate:        self.EndDate,
		Expenses:       self.Expenses,
		Active:         self.Active,
	}, nil
}

func (self *Budget) SetBSON(raw bson.Raw) error {
	type Alias Budget

	decoded := new(struct {
		ID             bson.ObjectId     `json:"id" bson:"_id"`
		InitialBalance float64           `json:"initialbalance" bson:"initialbalance"`
		StartDate      time.Time         `json:"startdate" bson:"startdate"`
		EndDate        time.Time         `json:"enddate" bson:"enddate"`
		Expenses       []expense.Expense `json:"expenses" bson:"expenses"`
		Active         bool              `json:"active" bson:"active"`
	})

	if err := raw.Unmarshal(&decoded); err != nil {
		return nil
	}

	self.ID = decoded.ID
	self.InitialBalance = decimal.NewFromFloat(decoded.InitialBalance)
	self.StartDate = decoded.StartDate
	self.EndDate = decoded.EndDate
	self.Expenses = decoded.Expenses
	self.Active = decoded.Active

	return nil
}

func NewBudget(balance decimal.Decimal, initialExpenses []expense.Expense) *Budget {
	instance := new(Budget)
	instance.StartDate = time.Now()
	instance.Active = true
	instance.Expenses = initialExpenses
	instance.InitialBalance = balance
	return instance
}
