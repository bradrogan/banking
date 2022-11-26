package app

import (
	"log"
	"net/http"

	"github.com/bradrogan/banking/config"
	"github.com/bradrogan/banking/connections"
	"github.com/bradrogan/banking/domain"
	"github.com/bradrogan/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	db := connections.NewDbClient()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb(db))}
	ah := AccountHandlers{service.NewAccountService(domain.NewAccountRepositoryDb(db))}

	router.HandleFunc("/customers", ch.getCustomersByStatus).Queries("status", "{status:[0-9a-zA-Z_]*}").Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]*}", ch.getCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customer_id:[0-9]*}/account", ah.NewAccount).Methods(http.MethodPost)

	addr := config.App.Server.Host + ":" + config.App.Server.Port
	log.Fatal(http.ListenAndServe(addr, router))
}
