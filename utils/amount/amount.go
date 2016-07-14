package amount

import (
	"errors"
	"strconv"
	"strings"
)

func ParseString(amount string) (float32, error) {
	amount = strings.Replace(amount, ",", ".", 1)
	amountFloat, err := strconv.ParseFloat(amount, 32)

	if err != nil {
		return -1, errors.New("The given string " + amount + " is not a number.")
	}

	return float32(amountFloat), nil
}
