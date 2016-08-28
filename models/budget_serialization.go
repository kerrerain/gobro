package models

import (
	"github.com/shopspring/decimal"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (self *Budget) GetBSON() (interface{}, error) {
	initialBalanceFloat, _ := self.InitialBalance.Float64()

	return &struct {
		InitialBalance   float64   `json:"initialbalance" bson:"initialbalance"`
		StartDate        time.Time `json:"startdate" bson:"startdate"`
		LastModification time.Time `json:"enddate" bson:"enddate"`
		Expenses         []Expense `json:"expenses" bson:"expenses"`
		Active           bool      `json:"active" bson:"active"`
	}{
		InitialBalance:   initialBalanceFloat,
		StartDate:        self.StartDate,
		LastModification: self.LastModification,
		Expenses:         self.Expenses,
		Active:           self.Active,
	}, nil
}

func (self *Budget) SetBSON(raw bson.Raw) error {
	type Alias Budget

	decoded := new(struct {
		ID               bson.ObjectId `json:"id" bson:"_id"`
		InitialBalance   float64       `json:"initialbalance" bson:"initialbalance"`
		StartDate        time.Time     `json:"startdate" bson:"startdate"`
		LastModification time.Time     `json:"enddate" bson:"enddate"`
		Expenses         []Expense     `json:"expenses" bson:"expenses"`
		Active           bool          `json:"active" bson:"active"`
	})

	if err := raw.Unmarshal(&decoded); err != nil {
		return nil
	}

	self.ID = decoded.ID
	self.InitialBalance = decimal.NewFromFloat(decoded.InitialBalance)
	self.StartDate = decoded.StartDate
	self.LastModification = decoded.LastModification
	self.Expenses = decoded.Expenses
	self.Active = decoded.Active

	return nil
}
