package dto

import (
	"strconv"

	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/helpers"
)

const MinimumOpeningBalance = 5000.00

type NewAccountRequest struct {
	Type       string  `json:"account_type,omitempty" xml:"account_type"`
	Amount     float64 `json:"amount,omitempty" xml:"amount"`
	CustomerId string  `json:"customer_id,omitempty" xml:"customer_id"`
}

var accountTypes = [...]string{
	"saving",
	"checking",
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if !helpers.Contains(r.Type, accountTypes[:]) {
		return errs.NewValidationError("invalid account type")
	}

	if r.Amount < MinimumOpeningBalance {
		return errs.NewValidationError("account opening requires a minimum $" + strconv.FormatFloat(MinimumOpeningBalance, 'f', 2, 64) + " deposit")
	}

	return nil
}
