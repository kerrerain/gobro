package budget_test

import (
	budgetPackage "github.com/magleff/gobro/features/budget"
	"github.com/magleff/gobro/features/expense"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewBudgetController(t *testing.T) {
	// Arrange
	// Act
	controller := budgetPackage.NewBudgetController().(*budgetPackage.BudgetControllerImpl)

	// Assert
	assert.NotNil(t, controller.BudgetDatastore, "Should initialize BudgetDatastore.")
	assert.NotNil(t, controller.ExpenseFixedDatastore, "Should initialize ExpenseFixedDatastore.")
}

func TestCreateBudget(t *testing.T) {
	// Arrange
	expenses := []expense.Expense{}
	balance := "100.25"
	parsedBalance := decimal.NewFromFloat(100.25)

	budgetDatastore := new(MockBudgetDatastore)
	budgetDatastore.On("CurrentBudget").Return(nil)
	budgetDatastore.On("CreateBudget", parsedBalance, expenses).Return()

	controller := new(budgetPackage.BudgetControllerImpl)
	controller.BudgetDatastore = budgetDatastore

	// Act
	err := controller.CreateBudget(balance, expenses)

	// Assert
	budgetDatastore.AssertExpectations(t)
	assert.NoError(t, err, "Should not throw an error.")
}

func TestCreateBudgetBadBalance(t *testing.T) {
	// Arrange
	expenses := []expense.Expense{}
	balance := "0az15e"

	controller := new(budgetPackage.BudgetControllerImpl)

	// Act
	err := controller.CreateBudget(balance, expenses)

	// Assert
	assert.Error(t, err, "Should throw an error if the balance is invalid.")
}

func TestCreateBudgetAlreadyAnActiveBudget(t *testing.T) {
	// Arrange
	expenses := []expense.Expense{}
	balance := "100.25"

	budgetDatastore := new(MockBudgetDatastore)
	budgetDatastore.On("CurrentBudget").Return(&budgetPackage.Budget{})

	controller := new(budgetPackage.BudgetControllerImpl)
	controller.BudgetDatastore = budgetDatastore

	// Act
	err := controller.CreateBudget(balance, expenses)

	// Assert
	budgetDatastore.AssertExpectations(t)
	assert.Error(t, err, "Should throw an error if already an active budget.")
}

func TestSaveBudget(t *testing.T) {
	// Arrange
	budget := &budgetPackage.Budget{}

	budgetDatastore := new(MockBudgetDatastore)
	budgetDatastore.On("Save", budget).Return()

	controller := new(budgetPackage.BudgetControllerImpl)
	controller.BudgetDatastore = budgetDatastore

	// Act
	controller.SaveBudget(budget)

	// Assert
	budgetDatastore.AssertExpectations(t)
}

func TestCurrentBudget(t *testing.T) {
	// Arrange
	budgetDatastore := new(MockBudgetDatastore)
	budgetDatastore.On("CurrentBudget").Return(&budgetPackage.Budget{})

	controller := new(budgetPackage.BudgetControllerImpl)
	controller.BudgetDatastore = budgetDatastore

	// Act
	controller.CurrentBudget()

	// Assert
	budgetDatastore.AssertExpectations(t)
}

func TestAddExpenseToCurrentBudget(t *testing.T) {
	// Arrange
	amount := "100.25"
	description := "Hotel"
	currentBudget := &budgetPackage.Budget{}
	currentBudget.Expenses = []expense.Expense{}

	budgetDatastore := new(MockBudgetDatastore)
	budgetDatastore.On("CurrentBudget").Return(currentBudget)
	budgetDatastore.On("Save", currentBudget).Return()

	controller := new(budgetPackage.BudgetControllerImpl)
	controller.BudgetDatastore = budgetDatastore

	// Act
	err := controller.AddExpenseToCurrentBudget(amount, description)

	// Assert
	budgetDatastore.AssertExpectations(t)
	assert.Equal(t, 1, len(currentBudget.Expenses),
		"Should add a new expense to the budget.")
	assert.Equal(t, decimal.NewFromFloat(-100.25), currentBudget.Expenses[0].Amount,
		"Should add a new expense with the given amount (transformed automatically to a negative value).")
	assert.Equal(t, description, currentBudget.Expenses[0].Description,
		"Should add a new expense with the given description.")
	assert.NoError(t, err, "Should not throw an error.")
}

func TestAddExpenseToCurrentBudgetNil(t *testing.T) {
	// Arrange
	amount := "100.25"
	description := "Hotel"

	budgetDatastore := new(MockBudgetDatastore)
	budgetDatastore.On("CurrentBudget").Return(nil)

	controller := new(budgetPackage.BudgetControllerImpl)
	controller.BudgetDatastore = budgetDatastore

	// Act
	err := controller.AddExpenseToCurrentBudget(amount, description)

	// Assert
	budgetDatastore.AssertExpectations(t)
	assert.Error(t, err, "Should fail if there is not any active budget.")
}

func TestAddRawExpensesToCurrentBudget(t *testing.T) {
	// Arrange
	expenses := []expense.Expense{
		{Amount: decimal.NewFromFloat(-25.30)},
		{Amount: decimal.NewFromFloat(-27.28)},
	}

	currentBudget := &budgetPackage.Budget{}
	currentBudget.Expenses = []expense.Expense{}

	budgetDatastore := new(MockBudgetDatastore)
	budgetDatastore.On("CurrentBudget").Return(currentBudget)
	budgetDatastore.On("Save", currentBudget).Return()

	controller := new(budgetPackage.BudgetControllerImpl)
	controller.BudgetDatastore = budgetDatastore

	// Act
	err := controller.AddRawExpensesToCurrentBudget(expenses)

	// Assert
	budgetDatastore.AssertExpectations(t)
	assert.Equal(t, 2, len(currentBudget.Expenses),
		"Should add the expenses to the budget.")
	assert.NoError(t, err, "Should not throw an error.")
}

func TestAddRawExpensesToCurrentBudgetNil(t *testing.T) {
	// Arrange
	expenses := []expense.Expense{
		{Amount: decimal.NewFromFloat(-25.30)},
		{Amount: decimal.NewFromFloat(-27.28)},
	}

	budgetDatastore := new(MockBudgetDatastore)
	budgetDatastore.On("CurrentBudget").Return(nil)

	controller := new(budgetPackage.BudgetControllerImpl)
	controller.BudgetDatastore = budgetDatastore

	// Act
	err := controller.AddRawExpensesToCurrentBudget(expenses)

	// Assert
	budgetDatastore.AssertExpectations(t)
	assert.Error(t, err, "Should fail if there is not any active budget.")
}

func TestCloseCurrentBudget(t *testing.T) {
	// Arrange
	currentBudget := &budgetPackage.Budget{}
	currentBudget.Active = true

	budgetDatastore := new(MockBudgetDatastore)
	budgetDatastore.On("CurrentBudget").Return(currentBudget)
	budgetDatastore.On("Save", currentBudget).Return()

	controller := new(budgetPackage.BudgetControllerImpl)
	controller.BudgetDatastore = budgetDatastore

	// Act
	err := controller.CloseCurrentBudget()

	// Assert
	budgetDatastore.AssertExpectations(t)

	assert.Equal(t, false, currentBudget.Active, "Should deactivate the current budget.")
	assert.Equal(t, time.Now().Format("2006-01-02"),
		currentBudget.EndDate.Format("2006-01-02"), "Should set a end date for the current budget.")
	assert.NoError(t, err, "Should not throw an error.")
}

func TestCloseCurrentBudgetNil(t *testing.T) {
	// Arrange
	budgetDatastore := new(MockBudgetDatastore)
	budgetDatastore.On("CurrentBudget").Return(nil)

	controller := new(budgetPackage.BudgetControllerImpl)
	controller.BudgetDatastore = budgetDatastore

	// Act
	err := controller.CloseCurrentBudget()

	// Assert
	budgetDatastore.AssertExpectations(t)
	assert.Error(t, err, "Should fail if there is not any active budget.")
}

func TestComputeBudgetInfo(t *testing.T) {
	// Arrange
	currentBudget := &budgetPackage.Budget{}
	currentBudget.InitialBalance = decimal.NewFromFloat(1114.25)
	currentBudget.Expenses = []expense.Expense{
		expense.Expense{Amount: decimal.NewFromFloat(20.51), Checked: false},
		expense.Expense{Amount: decimal.NewFromFloat(-30.68), Checked: true},
		expense.Expense{Amount: decimal.NewFromFloat(10.05), Checked: false},
		expense.Expense{Amount: decimal.NewFromFloat(-18.36), Checked: false},
	}

	budgetDatastore := new(MockBudgetDatastore)
	budgetDatastore.On("CurrentBudget").Return(currentBudget)

	controller := new(budgetPackage.BudgetControllerImpl)
	controller.BudgetDatastore = budgetDatastore

	// Act
	budgetInfo, err := controller.ComputeBudgetInfo()

	// Assert
	budgetDatastore.AssertExpectations(t)
	assert.Equal(t, decimal.NewFromFloat(-49.04), budgetInfo.TotalExpenses,
		"Should compute the total of expenses.")
	assert.Equal(t, decimal.NewFromFloat(30.56), budgetInfo.TotalEarnings,
		"Should compute the total of earnings.")
	assert.Equal(t, decimal.NewFromFloat(-18.36), budgetInfo.TotalUncheckedExpenses,
		"Should compute the total of unchecked expenses.")
	assert.Equal(t, decimal.NewFromFloat(1114.25), budgetInfo.InitialBalance,
		"Should copy the initial balance.")
	assert.Equal(t, currentBudget.StartDate, budgetInfo.StartDate,
		"Should copy the start date.")
	assert.Equal(t, decimal.NewFromFloat(-18.48), budgetInfo.Difference,
		"Should compute the difference.")
	assert.Equal(t, decimal.NewFromFloat(1095.77), budgetInfo.CurrentBalance,
		"Should compute the current balance.")
	assert.NoError(t, err, "Should not throw an error.")
}
