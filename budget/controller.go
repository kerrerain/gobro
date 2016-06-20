package budget

import (
	"gopkg.in/mgo.v2"
)

type BudgetController struct {
	session *mgo.Session
}

func Controller(session *mgo.Session) *BudgetController {
	return &BudgetController{session}
}

func (ec BudgetController) CreateBudget() {
	DataStore(ec.session).CreateBudget()
}
