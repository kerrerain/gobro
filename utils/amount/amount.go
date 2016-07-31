package amount

import (
	"errors"
	"github.com/shopspring/decimal"
	"strings"
)

// Creates a number from a string input, with 2 decimals.
// Example: 12.055 is replaced by 12.05 (the last digit was truncated).
func ParseString(amount string) (decimal.Decimal, error) {
	amount = strings.Replace(amount, ",", ".", 1)
	amountDecimal, err := decimal.NewFromString(amount)

	if err != nil {
		return decimal.NewFromFloat(-1),
			errors.New("The given string " + amount + " is not a number.")
	}

	return amountDecimal.Truncate(2), nil
}
