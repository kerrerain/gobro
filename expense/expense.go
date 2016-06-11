package expense

import (
	"time"
)

type Expense struct {
	Date        time.Time
	Description string
	Amount      float32
}
