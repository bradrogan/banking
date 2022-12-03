package dto

import (
	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/helpers"
)

type NewTransactionRequst struct {
	CustomerId string
	AccountId  string
	Type       string
	Amount     float64
}

var transactionTypes = [...]string{
	"deposit",
	"withdrawal",
}

func (t NewTransactionRequst) Validate() *errs.AppError {
	if !helpers.Contains(t.Type, transactionTypes[:]) {
		return errs.NewValidationError("invalid transaction type")
	}
	return nil
}
