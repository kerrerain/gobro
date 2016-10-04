package dao

import (
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/entities"
	"gopkg.in/mgo.v2/bson"
)

type BudgetDao interface {
	// Generic
	Create(budget entities.Budget) error
	Update(budget entities.Budget) error
	Delete(budget entities.Budget) error
	FindById(budgetId bson.ObjectId) (*entities.Budget, error)
	// Specific
	FindActiveBudget(accountId bson.ObjectId) (*entities.Budget, error)
}

type BudgetDaoImpl struct{}

func (e BudgetDaoImpl) Create(budget entities.Budget) error {
	var err error
	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("budget").Insert(budget)
	})
	return err
}

func (e BudgetDaoImpl) Update(budget entities.Budget) error {
	var err error
	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("budget").UpdateId(budget.ID, budget)
	})
	return err
}

func (e BudgetDaoImpl) Delete(budget entities.Budget) error {
	var err error
	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("budget").RemoveId(budget.ID)
	})
	return err
}

func (e BudgetDaoImpl) FindById(budgetId bson.ObjectId) (*entities.Budget, error) {
	var budget entities.Budget
	var err error

	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("budget").FindId(budgetId).One(&budget)
	})

	return &budget, err
}

// Given an account, looks for a budget whose end date is not set (currently active)
func (e BudgetDaoImpl) FindActiveBudget(accountId bson.ObjectId) (*entities.Budget, error) {
	var budget entities.Budget
	var err error

	database.ExecuteInSession(func(session database.Session) {
		err = session.DefaultSchema().Collection("budget").
			Find(bson.M{
				"enddate":   bson.M{"$exists": false},
				"accountid": accountId,
			}).One(&budget)
	})

	return &budget, err
}
