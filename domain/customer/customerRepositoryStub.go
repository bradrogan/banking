package customer

import "github.com/bradrogan/banking/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) ByActive(st Status) ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) ById(id string) (*Customer, *errs.AppError) {
	return &s.customers[0], nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id:          "1001",
			Name:        "Fred",
			City:        "Edmonton",
			Zipcode:     "90210",
			DateOfBirth: "Jan 2, 1972",
			Status:      1,
		},
		{
			Id:          "1002",
			Name:        "George",
			City:        "Calgary",
			Zipcode:     "78910",
			DateOfBirth: "Jun 4, 1988",
			Status:      1,
		},
	}
	return CustomerRepositoryStub{customers: customers}
}
