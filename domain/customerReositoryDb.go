package domain

import (
	"database/sql"
	"time"

	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type customerRepositoryDb struct {
	client *sqlx.DB
}

const (
	DbCustomerActive   = 1
	DbCustomerInactive = 0
)

func (d customerRepositoryDb) ByActive(status CustomerStatus) ([]Customer, *errs.AppError) {
	byActiveSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
	var byActiveQueryParam uint

	customers := make([]Customer, 0)

	switch status {
	case CustomerStatusActive:
		byActiveQueryParam = DbCustomerActive
	case CustomerStatusInactive:
		byActiveQueryParam = DbCustomerInactive
	}

	err := d.client.Select(&customers, byActiveSql, byActiveQueryParam)

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error: " + err.Error())
	}

	return customers, nil
}

func (d customerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	customers := make([]Customer, 0)

	err := d.client.Select(&customers, findAllSql)

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error: " + err.Error())
	}

	return customers, nil
}

func (d customerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scanning customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error" + err.Error())
	}
	return &c, nil
}

func NewCustomerRepositoryDb() customerRepositoryDb {
	client, err := sqlx.Open("mysql", "root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return customerRepositoryDb{client: client}
}
