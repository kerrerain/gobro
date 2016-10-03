package models

import (
	"github.com/magleff/gobro/database"
	"github.com/shopspring/decimal"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type BudgetEntity interface {
	FindById(budgetId bson.ObjectId) (*Budget, error)
}

type Budget struct {
	ID               bson.ObjectId `bson:"_id,omitempty"`
	AccountId        bson.ObjectId `bson:"accountid,omitempty"`
	UserId           bson.ObjectId `bson:"userid,omitempty"`
	StartDate        time.Time
	LastModification time.Time
	Expenses         []Expense
	InitialBalance   decimal.Decimal
	Active           bool
}

func (e Budget) FindById(budgetId bson.ObjectId) (*Budget, error) {
	var budget Budget
	var err error

	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("budget").FindId(budgetId).One(&budget)
	})

	return &budget, err
}
