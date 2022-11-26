package domain

import (
	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
)

type Account struct {
	Id          string `db:"account_id"`
	Type        string `db:"account_type"`
	Amount      float64
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	Status      AccountStatus
}

type AccountStatus uint

const (
	AccountStatusInactive AccountStatus = iota
	AccountStatusActive
	endAccountStatus
)

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (status AccountStatus) IsValid() bool {
	return status < endAccountStatus
}

func (status AccountStatus) StatusAsText() string {
	switch status {
	case AccountStatusInactive:
		return "inactive"
	case AccountStatusActive:
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
