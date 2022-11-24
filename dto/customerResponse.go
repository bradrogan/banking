package dto

type CustomerStatus string

type CustomerResponse struct {
	Id          string `json:"id,omitempty" xml:"id"`
	Name        string `json:"name,omitempty" xml:"name"`
	City        string `json:"city,omitempty" xml:"city"`
	Zipcode     string `json:"zipcode,omitempty" xml:"zipcode"`
	DateOfBirth string `json:"date_of_birth,omitempty" xml:"date_of_birth"`
	Status      string `json:"status,omitempty" xml:"status"`
}
