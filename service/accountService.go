package service

import (
	"time"

	"github.com/bradrogan/banking/domain"
	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/logger"
)

type AccountService struct {
	repo domain.AccountRepository
}

func (s AccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		logger.Error("validaton failed")
		return nil, err
	}
	a := domain.Account{
		Id:          "",
		Type:        domain.AccountType(req.Type),
		Amount:      req.Amount,
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		Status:      domain.AccountActive,
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
