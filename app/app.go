package app

import (
	"log"
	"net/http"

	"github.com/bradrogan/banking/config"
	"github.com/bradrogan/banking/domain"
	"github.com/bradrogan/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getCustomersByStatus).Queries("status", "{status:[0-9a-zA-Z_]*}").Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]*}", ch.getCustomer).Methods(http.MethodGet)

	addr := config.App.Server.Host + ":" + config.App.Server.Port
	log.Fatal(http.ListenAndServe(addr, router))
}
