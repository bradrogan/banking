package app

import (
	"net/http"

	"github.com/bradrogan/banking/domain"
	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
	"github.com/gorilla/mux"
)

type CustomerServicer interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomersByStatus(domain.CustomerStatus) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}
type CustomerHandlers struct {
	service CustomerServicer
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomers()

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, customers)
}

func (ch *CustomerHandlers) getCustomersByStatus(w http.ResponseWriter, r *http.Request) {
	activeParam := mux.Vars(r)["status"]
	var status domain.CustomerStatus

	switch activeParam {
	case domain.CustomerStatusActive.StatusAsText():
		status = domain.CustomerStatusActive
	case domain.CustomerStatusInactive.StatusAsText():
		status = domain.CustomerStatusInactive
	default:
		writeResponse(w, http.StatusBadRequest, "invalid value for 'status' query parameter")
		return
	}

	customers, err := ch.service.GetCustomersByStatus(status)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, customers)

}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, customer)
}
