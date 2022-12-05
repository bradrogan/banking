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
	GetAccount(string) (*account.Account, *errs.AppError)
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
func (s AccountService) GetAccount(id string) (*dto.AccountResponse, *errs.AppError) {
	a, err := s.data.GetAccount(id)
	if err != nil {
		return nil, err
	}
	response := a.ToAccountDto()

	return &response, nil
}

func New(a accountData) *AccountService {
	return &AccountService{
		data: a,
	}
}
