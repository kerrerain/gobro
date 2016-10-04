package entities

import (
	"github.com/shopspring/decimal"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (self *Budget) GetBSON() (interface{}, error) {
	initialBalanceFloat, _ := self.InitialBalance.Float64()

	return &struct {
		AccountId        bson.ObjectId `bson:"accountid,omitempty"`
		UserId           bson.ObjectId `bson:"userid,omitempty"`
		InitialBalance   float64       `json:"initialbalance" bson:"initialbalance"`
		StartDate        time.Time     `json:"startdate" bson:"startdate"`
		EndDate          time.Time     `json:"enddate,omitempty" bson:"enddate,omitempty"`
		LastModification time.Time     `json:"lastmodification" bson:"lastmodification"`
		Expenses         []Expense     `json:"expenses" bson:"expenses"`
		Active           bool          `json:"active" bson:"active"`
	}{
		AccountId:        self.AccountId,
		UserId:           self.UserId,
		InitialBalance:   initialBalanceFloat,
		StartDate:        self.StartDate,
		EndDate:          self.EndDate,
		LastModification: self.LastModification,
		Expenses:         self.Expenses,
		Active:           self.Active,
	}, nil
}

func (self *Budget) SetBSON(raw bson.Raw) error {
	decoded := new(struct {
		ID               bson.ObjectId `json:"id" bson:"_id"`
		AccountId        bson.ObjectId `bson:"accountid,omitempty"`
		UserId           bson.ObjectId `bson:"userid,omitempty"`
		InitialBalance   float64       `json:"initialbalance" bson:"initialbalance"`
		StartDate        time.Time     `json:"startdate" bson:"startdate"`
		EndDate          time.Time     `json:"enddate,omitempty" bson:"enddate,omitempty"`
		LastModification time.Time     `json:"lastmodification" bson:"lastmodification"`
		Expenses         []Expense     `json:"expenses" bson:"expenses"`
		Active           bool          `json:"active" bson:"active"`
	})

	if err := raw.Unmarshal(&decoded); err != nil {
		return nil
	}

	self.ID = decoded.ID
	self.AccountId = decoded.AccountId
	self.UserId = decoded.UserId
	self.InitialBalance = decimal.NewFromFloat(decoded.InitialBalance)
	self.StartDate = decoded.StartDate
	self.EndDate = decoded.EndDate
	self.LastModification = decoded.LastModification
	self.Expenses = decoded.Expenses
	self.Active = decoded.Active

	return nil
}
