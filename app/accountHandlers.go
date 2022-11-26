package app

import (
	"net/http"

	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
	"github.com/gorilla/mux"
)

type AccountServicer interface {
	NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type AccountHandlers struct {
	service AccountServicer
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
