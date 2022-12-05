package app

import (
	"net/http"

	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
	"github.com/gorilla/mux"
)

type transactionSaver interface {
	SaveTransaction(dto.NewTransactionRequst) (*dto.NewTransactionResponse, *errs.AppError)
}

type TransactionHandlers struct {
	transactionSaver transactionSaver
}

func (th *TransactionHandlers) NewTransaction(w http.ResponseWriter, r *http.Request) {
	var request dto.NewTransactionRequst

	vars := mux.Vars(r)
	customer := vars["customer_id"]
	account := vars["account_id"]

	request.CustomerId = customer
	request.AccountId = account

	err := readRequest(r, &request)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	response, err := th.transactionSaver.SaveTransaction(request)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, response)
}

func NewTransactionHandler(t transactionSaver) *TransactionHandlers {
	return &TransactionHandlers{
		transactionSaver: t,
	}
}
