package models

import (
	"fmt"
)

func (e Budget) Display() {
	totalEarnings := ComputeTotalEarnings(e.Expenses)
	totalExpenses := ComputeTotalExpenses(e.Expenses)
	totalUncheckedExpenses := ComputeTotalUncheckedExpenses(e.Expenses)
	difference := totalEarnings.Add(totalExpenses)
	balance := e.InitialBalance.Add(difference)

	fmt.Println("Created on", e.StartDate)
	fmt.Println("Initial balance", e.InitialBalance)
	fmt.Println("Total earnings", totalEarnings)
	fmt.Println("Total expenses", totalExpenses)
	fmt.Println("Total unchecked expenses", totalUncheckedExpenses)
	fmt.Println("Balance", balance.String(), "("+difference.String()+")")
}
