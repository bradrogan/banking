package service

import (
	"time"

	"github.com/bradrogan/banking/domain"
	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
)

type AccountService struct {
	repo domain.AccountRepository
}

func (s AccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	a := domain.Account{
		Id:          "",
		Type:        req.Type,
		Amount:      req.Amount,
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		Status:      domain.AccountStatusActive,
	}

	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToDto()

	return &response, nil

}

func NewAccountService(repository domain.AccountRepository) AccountService {
	return AccountService{repository}
}
