package transactionsvc

import (
	"time"

	"github.com/bradrogan/banking/domain/account"
	"github.com/bradrogan/banking/domain/transaction"
	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/logger"
)

type transactionData interface {
	Save(transaction.Transaction) (*transaction.Transaction, *errs.AppError)
}

type accountData interface {
	GetAccount(string) (*account.Account, *errs.AppError)
	UpdateBalance(account.Account) (*account.Account, *errs.AppError)
}

type TransactionService struct {
	transactionData transactionData
	accountData     accountData
}

func (ts TransactionService) SaveTransaction(req dto.NewTransactionRequst) (*dto.NewTransactionResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		logger.Error("transaction validation failed")
		return nil, err
	}

	t := transaction.Transaction{
		AccountId: req.AccountId,
		Type:      transaction.Type(req.Type),
		Amount:    req.Amount,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
	}

	a, err := ts.accountData.GetAccount(t.AccountId)
	if err != nil {
		return nil, err
	}

	verr := verifyTransaction(a, req.CustomerId, t.Amount, t.Type)
	if verr != nil {
		return nil, verr
	}

	newTransaction, err := ts.transactionData.Save(t)
	if err != nil {
		return nil, err
	}

	switch t.Type {
	case transaction.Deposit:
		a.Amount += t.Amount
	case transaction.Withdrawal:
		a.Amount -= t.Amount
	}

	ts.accountData.UpdateBalance(*a)

	response := dto.NewTransactionResponse{
		Id:      newTransaction.Id,
		Balance: 0.0,
	}

	return &response, nil
}

func verifyTransaction(a *account.Account, customerId string, amount float64, t transaction.Type) *errs.AppError {
	if a.CustomerId != customerId {
		return errs.NewValidationError("customer id does not match account holder")
	}

	if t == transaction.Withdrawal && amount > a.Amount {
		return errs.NewValidationError("withdrawal amount exceeds account balance")
	}

	return nil
}

func New(td transactionData, ad accountData) *TransactionService {
	return &TransactionService{
		transactionData: td,
		accountData:     ad,
	}
}
