package customersvc

import (
	"github.com/bradrogan/banking/domain/customer"
	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
)

//go:generate mockgen -destination=../../mock/mockCustomerData.go -package=mock -source=./customerSvc.go customerData
type customerData interface {
	FindAll() ([]customer.Customer, *errs.AppError)
	ById(string) (*customer.Customer, *errs.AppError)
	ByActive(customer.Status) ([]customer.Customer, *errs.AppError)
}

type CustomerService struct {
	data customerData
}

func (s CustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	c, err := s.data.FindAll()
	if err != nil {
		return nil, err
	}
	return customer.ToDto(c), nil
}

func (s CustomerService) GetCustomersByStatus(status customer.Status) ([]dto.CustomerResponse, *errs.AppError) {
	c, err := s.data.ByActive(status)
	if err != nil {
		return nil, err
	}
	return customer.ToDto(c), nil
}

func (s CustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.data.ById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil

}

func New(cd customerData) CustomerService {
	return CustomerService{data: cd}
}
