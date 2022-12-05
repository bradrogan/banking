package app

import (
	"net/http"

	"github.com/bradrogan/banking/domain/customer"
	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
	"github.com/gorilla/mux"
)

type CustomerServicer interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomersByStatus(customer.Status) ([]dto.CustomerResponse, *errs.AppError)
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
	var status customer.Status

	switch activeParam {
	case customer.Active.StatusAsText():
		status = customer.Active
	case customer.Inactive.StatusAsText():
		status = customer.Inactive
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
