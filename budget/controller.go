package budget

import (
	"github.com/magleff/gobro/database"
)

type BudgetController struct {
	Datastore *BudgetDatastore
}

func NewController(DB *database.Database) *BudgetController {
	instance := new(BudgetController)
	instance.Datastore = NewDatastore(DB)
	return instance
}

func (self BudgetController) CreateBudget() {
	self.Datastore.CreateBudget()
}
