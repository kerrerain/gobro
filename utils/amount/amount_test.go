package amount

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInvalidString(t *testing.T) {
	str := "aAnjbssvg"
	_, err := ParseString(str)
	assert.EqualError(t, err, "The given string "+str+" is not a number.",
		"An invalid string should raise an error.")
}

func TestParseCommaString(t *testing.T) {
	str := "0,3"
	amount, _ := ParseString(str)
	assert.Equal(t, decimal.NewFromFloat(0.3), amount, "Should parse a string even with a comma.")
}

func TestParseString(t *testing.T) {
	str := "0.345"
	amount, _ := ParseString(str)
	assert.Equal(t, decimal.NewFromFloat(0.34), amount, "Should parse a string and truncate to 2 decimals.")
}
