package expensefixed

import (
	"time"
)

type ExpenseFixed struct {
	Date        time.Time
	Description string
	Amount      float32
}
