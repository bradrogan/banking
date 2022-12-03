package service

import (
	"time"

	"github.com/bradrogan/banking/domain"
	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/logger"
)

type saver interface {
	Save(domain.Transaction) (*domain.Transaction, *errs.AppError)
}

type customerGetter interface {
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type AccountGetter interface {
	//get account by ID
}

type TransactionService struct {
	saver saver
}

func (ts TransactionService) SaveTransaction(req dto.NewTransactionRequst) (*dto.NewTransactionResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		logger.Error("transaction validation failed")
		return nil, err
	}
	//check customer

	//check balance

	//save transaction
	t := domain.Transaction{
		AccountId: req.AccountId,
		Type:      domain.TransactionType(req.Type),
		Amount:    req.Amount,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
	}

	newTransaction, err := ts.saver.Save(t)
	if err != nil {
		return nil, err
	}

	response := dto.NewTransactionResponse{
		Id:      newTransaction.Id,
		Balance: 0.0,
	}

	return &response, nil
}

func NewTransactionService(s saver) *TransactionService {
	return &TransactionService{
		saver: s,
	}
}
