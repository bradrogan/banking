package app

import (
	"log"
	"net/http"

	"github.com/bradrogan/banking/config"
	"github.com/bradrogan/banking/connections"
	"github.com/bradrogan/banking/domain/account"
	"github.com/bradrogan/banking/domain/customer"
	"github.com/bradrogan/banking/domain/transaction"
	"github.com/bradrogan/banking/service/accountsvc"
	"github.com/bradrogan/banking/service/customersvc"
	"github.com/bradrogan/banking/service/transactionsvc"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	db := connections.NewDbClient()

	ch := CustomerHandlers{customersvc.New(customer.NewCustomerRepositoryDb(db))}

	ah := NewAccountHandler(
		accountsvc.New(
			account.NewDbRepository(db)))

	th := NewTransactionHandler(
		transactionsvc.New(
			transaction.NewDbRepository(db), account.NewDbRepository(db)))

	router.HandleFunc("/customers", ch.getCustomersByStatus).Queries("status", "{status:[0-9a-zA-Z_]*}").Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]*}", ch.getCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customer_id:[0-9]*}/account", ah.NewAccount).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]*}/account/{account_id:[0-9]*}/transaction", th.NewTransaction).Methods(http.MethodPost)

	addr := config.App.Server.Host + ":" + config.App.Server.Port
	log.Fatal(http.ListenAndServe(addr, router))
}
