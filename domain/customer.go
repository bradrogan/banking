package domain

import (
	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      CustomerStatus
}

type CustomerStatus uint

const (
	CustomerStatusInactive CustomerStatus = iota
	CustomerStatusActive
	end
)

func (c CustomerStatus) IsValid(value uint) bool {
	return value < uint(end)
}

func (status CustomerStatus) StatusAsText() string {
	switch status {
	case CustomerStatusInactive:
		return "inactive"
	case CustomerStatusActive:
		return "active"
	default:
		return ""
	}
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
	ByActive(CustomerStatus) ([]Customer, *errs.AppError)
}

func (c Customer) ToDto() dto.CustomerResponse {

	var statusText string

	switch c.Status {
	case DbCustomerActive:
		statusText = "active"
	case DbCustomerInactive:
		statusText = "inactive"
	}

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      statusText,
	}
}

func ToDto(c []Customer) []dto.CustomerResponse {
	response := make([]dto.CustomerResponse, 0)
	for _, val := range c {
		response = append(response, val.ToDto())
	}
	return response
}
