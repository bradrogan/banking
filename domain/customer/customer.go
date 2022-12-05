package customer

import (
	"github.com/bradrogan/banking/dto"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      Status
}

//go:generate stringer -type=Status
type Status uint

const (
	Inactive Status = iota
	Active
	end
)

func (cs Status) IsValid() bool {
	return cs < end
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

func (c Customer) ToDto() dto.CustomerResponse {

	var statusText string

	switch c.Status {
	case Active:
		statusText = "active"
	case Inactive:
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
