package models

import (
	"github.com/magleff/gobro/database"
	"github.com/shopspring/decimal"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type BudgetEntity interface {
	GetCurrent() *Budget
}

type Budget struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	StartDate      time.Time
	EndDate        time.Time
	Expenses       []Expense
	InitialBalance decimal.Decimal
	Active         bool
}

func (e Budget) GetCurrent() *Budget {
	var budget Budget
	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("budget").Find(bson.M{"active": true}).One(&budget)
	})
	return &budget
}
