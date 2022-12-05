package accountsvc

import (
	"time"

	"github.com/bradrogan/banking/domain/account"
	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/logger"
)

type accountData interface {
	Save(account.Account) (*account.Account, *errs.AppError)
	GetAccount(string, string) (*account.Account, *errs.AppError)
	SaveTransaction(account.Transaction) (*account.Transaction, *errs.AppError)
}

type AccountService struct {
	data accountData
}

func (s AccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		logger.Error("validaton failed")
		return nil, err
	}
	a := account.Account{
		Id:          "",
		Type:        account.Type(req.Type),
		Amount:      req.Amount,
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		Status:      account.Active,
	}

	newAccount, err := s.data.Save(a)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountDto()

	return &response, nil

}

func (s AccountService) SaveTransaction(req dto.NewTransactionRequst) (*dto.NewTransactionResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		logger.Error("transaction validation failed")
		return nil, err
	}

	a, err := s.data.GetAccount(req.AccountId, req.CustomerId)
	if err != nil {
		return nil, err
	}

	if req.IsWithdrawal() && !a.CanWithdraw(req.Amount) {
		return nil, errs.NewValidationError("withdrawal amount exceeds limit")
	}

	t := account.Transaction{
		AccountId: req.AccountId,
		Type:      account.TransactionType(req.Type),
		Amount:    req.Amount,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
	}

	newTransaction, err := s.data.SaveTransaction(t)
	if err != nil {
		return nil, err
	}

	a, _ = s.data.GetAccount(req.AccountId, req.CustomerId)

	response := dto.NewTransactionResponse{
		TransactionId:   newTransaction.Id,
		TransactionType: string(newTransaction.Type),
		AccountId:       newTransaction.AccountId,
		Balance:         a.Amount,
		TransactionDate: newTransaction.Time,
	}

	return &response, nil
}

func New(a accountData) *AccountService {
	return &AccountService{
		data: a,
	}
}
