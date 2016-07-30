package budget_test

import (
	budgetPackage "github.com/magleff/gobro/features/budget"
	"testing"
)

func TestCurrentBudget(t *testing.T) {
	// Arrange
	budgetDatastore := new(MockBudgetDatastore)
	budgetDatastore.On("CurrentBudget").Return(&budgetPackage.Budget{})

	controller := new(budgetPackage.BudgetControllerImpl)
	controller.BudgetDatastore = budgetDatastore

	// Act
	controller.CurrentBudget()

	// Assert
	budgetDatastore.AssertCalled(t, "CurrentBudget")
}

func TestSaveBudget(t *testing.T) {
	// Arrange
	budgetDatastore := new(MockBudgetDatastore)
	budgetDatastore.On("Save", &budgetPackage.Budget{}).Return()

	controller := new(budgetPackage.BudgetControllerImpl)
	controller.BudgetDatastore = budgetDatastore

	budget := &budgetPackage.Budget{}

	// Act
	controller.SaveBudget(budget)

	// Assert
	budgetDatastore.AssertCalled(t, "Save", budget)
}
