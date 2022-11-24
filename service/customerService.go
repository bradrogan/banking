package service

import (
	"github.com/bradrogan/banking/domain"
	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
)

type CustomerService struct {
	repo domain.CustomerRepository
}

func (s CustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return domain.ToDto(c), nil
}

func (s CustomerService) GetCustomersByStatus(status domain.CustomerStatus) ([]dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ByActive(status)
	if err != nil {
		return nil, err
	}
	return domain.ToDto(c), nil
}

func (s CustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil

}

func NewCustomerService(repository domain.CustomerRepository) CustomerService {
	return CustomerService{repo: repository}
}
