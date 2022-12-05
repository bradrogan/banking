package app

import (
	"net/http"

	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
	"github.com/gorilla/mux"
)

type accountServicer interface {
	NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	SaveTransaction(dto.NewTransactionRequst) (*dto.NewTransactionResponse, *errs.AppError)
}

type AccountHandlers struct {
	service accountServicer
}

func (ah *AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {

	var request dto.NewAccountRequest

	vars := mux.Vars(r)
	id := vars["customer_id"]
	request.CustomerId = id

	err := readRequest(r, &request)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	response, err := ah.service.NewAccount(request)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, response)
}

func (ah *AccountHandlers) NewTransaction(w http.ResponseWriter, r *http.Request) {
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

	response, err := ah.service.SaveTransaction(request)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, response)
}

func NewAccountHandler(a accountServicer) *AccountHandlers {
	return &AccountHandlers{
		service: a,
	}
}
