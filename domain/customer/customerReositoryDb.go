package customer

import (
	"database/sql"

	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type customerRepositoryDb struct {
	client *sqlx.DB
}

func (d customerRepositoryDb) ByActive(status Status) ([]Customer, *errs.AppError) {
	byActiveSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"

	if ok := status.IsValid(); !ok {
		logger.Error("invalid customer status enum value: ", zap.Uint("customer_status", uint(status)))
		return nil, errs.NewUnexpectedError("invalid customer status value")
	}

	customers := make([]Customer, 0)

	err := d.client.Select(&customers, byActiveSql, status)

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

	if err == sql.ErrNoRows {
		return nil, errs.NewNotFoundError("Customer not found")
	}
	if err != nil {
		logger.Error("Error while scanning customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error" + err.Error())
	}
	return &c, nil
}

func NewCustomerRepositoryDb(db *sqlx.DB) customerRepositoryDb {
	return customerRepositoryDb{client: db}
}
