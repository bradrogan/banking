package service

import (
	"github.com/bradrogan/banking/domain"
	"github.com/bradrogan/banking/errs"
)

type CustomerService struct {
	repo domain.CustomerRepository
}

func (s CustomerService) GetAllCustomers(status domain.CustomerStatus) ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll(status)
}

func (s CustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) CustomerService {
	return CustomerService{repo: repository}
}
