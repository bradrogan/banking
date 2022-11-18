package app

import (
	"log"
	"net/http"

	"github.com/bradrogan/banking/domain"
	"github.com/bradrogan/banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
