package dao

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/entities"
	"gopkg.in/mgo.v2/bson"
)

type BudgetDao interface {
	FindById(budgetId bson.ObjectId) (*entities.Budget, error)
}

type BudgetDaoImpl struct{}

func (e BudgetDaoImpl) FindById(budgetId bson.ObjectId) (*entities.Budget, error) {
	var budget entities.Budget
	var err error

	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("budget").FindId(budgetId).One(&budget)
	})

	return &budget, err
}
