package dao_test

import (
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/database"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

func TestFindActiveBudget(t *testing.T) {
	// Arrange
	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().DropDatabase()
	})

	accountId := bson.NewObjectId()

	budget1 := bson.M{
		"accountid": accountId,
		"enddate":   time.Now(),
		"active":    false,
	}

	budget2 := bson.M{
		"accountid": accountId,
		"active":    true,
	}

	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("budget").Insert(budget1, budget2)
	})

	budgetDao := dao.BudgetDaoImpl{}

	// Act
	activeBudget, err := budgetDao.FindActiveBudget(accountId)

	// Assert
	assert.NotNil(t, activeBudget, "Should find an active budget.")
	assert.NoError(t, err, "Should not raise an error.")
	assert.Equal(t, true, activeBudget.Active, "")
}

func TestFindActiveBudgetNotFound(t *testing.T) {
	// Arrange
	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().DropDatabase()
	})

	accountId := bson.NewObjectId()

	budget1 := bson.M{
		"accountid": accountId,
		"enddate":   time.Now(),
	}

	database.ExecuteInSession(func(session database.Session) {
		session.DefaultSchema().Collection("budget").Insert(budget1)
	})

	budgetDao := dao.BudgetDaoImpl{}

	// Act
	_, err := budgetDao.FindActiveBudget(accountId)

	// Assert
	assert.Error(t, err, "Should raise an error.")
}
