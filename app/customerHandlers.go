package app

import (
	"encoding/json"
	"net/http"

	"github.com/bradrogan/banking/domain"
	"github.com/bradrogan/banking/errs"
	"github.com/gorilla/mux"
)

type CustomerServicer interface {
	GetAllCustomers() ([]domain.Customer, *errs.AppError)
	GetCustomersByStatus(domain.CustomerStatus) ([]domain.Customer, *errs.AppError)
	GetCustomer(id string) (*domain.Customer, *errs.AppError)
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
	case "active":
		status = domain.CustomerStatusActive
	case "inactive":
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

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}

}
