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
	client *sql.DB
}

const (
	DbCustomerActive   = "1"
	DbCustomerInactive = "0"
)

func (d customerRepositoryDb) ByActive(status CustomerStatus) ([]Customer, *errs.AppError) {
	byActiveSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
	var byActiveQueryParam string

	switch status {
	case CustomerStatusActive:
		byActiveQueryParam = DbCustomerActive
	case CustomerStatusInactive:
		byActiveQueryParam = DbCustomerInactive
	}

	rows, err := d.client.Query(byActiveSql, byActiveQueryParam)

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error: " + err.Error())
	}

	return parseCustomerResults(rows)
}

func (d customerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSql)

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error: " + err.Error())
	}

	return parseCustomerResults(rows)
}

func parseCustomerResults(rows *sql.Rows) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	err := sqlx.StructScan(rows, &customers)

	if err != nil {
		logger.Error("Error while scanning customers " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error: " + err.Error())
	}
	return customers, nil
}

func NewCustomerRepositoryDb() customerRepositoryDb {
	client, err := sql.Open("mysql", "root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return customerRepositoryDb{client: client}
}

func (d customerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)

	var c Customer

	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scanning customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &c, nil
}
