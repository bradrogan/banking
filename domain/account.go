package domain

import (
	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
)

type Account struct {
	Id          string      `db:"account_id"`
	Type        AccountType `db:"account_type"`
	Amount      float64
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	Status      AccountStatus
}

type AccountStatus uint

const (
	AccountInactive AccountStatus = iota
	AccountActive
	endAccountStatus
)

type AccountType string

const (
	Checking AccountType = "checking"
	Saving   AccountType = "saving"
)

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (status AccountStatus) IsValid() bool {
	return status < endAccountStatus
}

func (status AccountStatus) StatusAsText() string {
	switch status {
	case AccountInactive:
		return "inactive"
	case AccountActive:
		return "active"
	default:
		return ""
	}
}

func (a Account) ToDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountID: a.Id,
	}
}
