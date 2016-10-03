package entities

import (
	"github.com/shopspring/decimal"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (self *Expense) GetBSON() (interface{}, error) {
	amountFloat, _ := self.Amount.Float64()

	// A custom definition of the struct is needed:
	// BSON can't marshal the Decimal type, Amount is switched to a float64.
	return &struct {
		Date        time.Time `json:"date" bson:"date"`
		Description string    `json:"description" bson:"description"`
		Amount      float64   `json:"amount" bson:"amount"`
		Checked     bool      `json:"checked" bson:"checked"`
	}{
		Date:        self.Date,
		Description: self.Description,
		Amount:      amountFloat,
		Checked:     self.Checked,
	}, nil
}

func (self *Expense) SetBSON(raw bson.Raw) error {
	decoded := new(struct {
		ID          bson.ObjectId `json:"_id" bson:"_id"`
		Date        time.Time     `json:"date" bson:"date"`
		Description string        `json:"description" bson:"description"`
		Amount      float64       `json:"amount" bson:"amount"`
		Checked     bool          `json:"checked" bson:"checked"`
	})

	if err := raw.Unmarshal(decoded); err != nil {
		return nil
	}

	self.Amount = decimal.NewFromFloat(decoded.Amount)
	self.ID = decoded.ID
	self.Date = decoded.Date
	self.Description = decoded.Description
	self.Checked = decoded.Checked

	return nil
}
