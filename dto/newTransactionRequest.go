package dto

import (
	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/helpers"
)

type NewTransactionRequst struct {
	CustomerId string  `json:"customer_id,omitempty" xml:"customer_id"`
	AccountId  string  `json:"account_id,omitempty" xml:"account_id"`
	Type       string  `json:"transaction_type,omitempty" xml:"type"`
	Amount     float64 `json:"amount,omitempty" xml:"amount"`
}

var transactionTypes = [...]string{
	"deposit",
	"withdrawal",
}

func (t NewTransactionRequst) Validate() *errs.AppError {
	if !helpers.Contains(t.Type, transactionTypes[:]) {
		return errs.NewValidationError("invalid transaction type")
	}

	if t.Amount < 0 {
		return errs.NewValidationError("amount cannot be negative")
	}

	return nil
}

func (t NewTransactionRequst) IsWithdrawal() bool {
	return t.Type == "withdrawal"
}
