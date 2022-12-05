package account

import "github.com/bradrogan/banking/dto"

type Account struct {
	Id          string `db:"account_id"`
	Type        Type   `db:"account_type"`
	Amount      float64
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	Status      Status
}

type Status uint

const (
	Inactive Status = iota
	Active
	end
)

type Type string

const (
	Checking Type = "checking"
	Saving   Type = "saving"
)

func (status Status) IsValid() bool {
	return status < end
}

func (status Status) StatusAsText() string {
	switch status {
	case Inactive:
		return "inactive"
	case Active:
		return "active"
	default:
		return ""
	}
}

func (a Account) ToNewAccountDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountID: a.Id,
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	return amount <= a.Amount
}
