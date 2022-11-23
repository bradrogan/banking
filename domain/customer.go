package domain

import "github.com/bradrogan/banking/errs"

type Customer struct {
	Id          string `json:"id,omitempty" xml:"id"`
	Name        string `json:"name,omitempty" xml:"name"`
	City        string `json:"city,omitempty" xml:"city"`
	Zipcode     string `json:"zipcode,omitempty" xml:"zipcode"`
	DateOfBirth string `json:"date_of_birth,omitempty" xml:"date_of_birth"`
	Status      string `json:"status,omitempty" xml:"status"`
}

type CustomerStatus uint

const (
	CustomerStatusAll CustomerStatus = iota
	CustomerStatusActive
	CustomerStatusInactive
)

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
	ByActive(CustomerStatus) ([]Customer, *errs.AppError)
}
