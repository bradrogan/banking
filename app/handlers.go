package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/bradrogan/banking/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomers()

	if err != nil {
		fmt.Fprint(w, "internal server error")
		w.WriteHeader(http.StatusInternalServerError)
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
