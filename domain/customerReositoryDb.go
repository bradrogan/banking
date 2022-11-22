package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/bradrogan/banking/errs"
	_ "github.com/go-sql-driver/mysql"
)

type customerRepositoryDb struct {
	client *sql.DB
}

const (
	DbCustomerActive   = "1"
	DbCustomerInactive = "0"
)

func (d customerRepositoryDb) FindAll(status CustomerStatus) ([]Customer, *errs.AppError) {

	var findAllSql string

	switch status {
	case CustomerStatusActive:
		findAllSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = " + DbCustomerActive
	case CustomerStatusInactive:
		findAllSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = " + DbCustomerInactive
	case CustomerStatusAll:
		findAllSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	}

	rows, err := d.client.Query(findAllSql)

	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error: " + err.Error())
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err = rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("no customers found: " + err.Error())
			}
			log.Println("Error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error: " + err.Error())
		}
		customers = append(customers, c)
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
		log.Println("Error while scanning customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &c, nil
}
