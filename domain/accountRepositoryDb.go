package domain

import (
	"strconv"

	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type accountRepositoryDb struct {
	client *sqlx.DB
}

func (d accountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	saveSql := "insert into accounts (type, amount, customer_id, opening_date, status) values (?, ?, ?, ?, ?)"

	res, err := d.client.Exec(saveSql, a.Type, a.Amount, a.CustomerId, a.OpeningDate, a.Status)
	if err != nil {
		logger.Error("error saving account", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Error("error getting last inserted account", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	a.Id = strconv.FormatInt(id, 10)

	return &a, nil

}
func NewAccountRepositoryDb(db *sqlx.DB) accountRepositoryDb {
	return accountRepositoryDb{client: db}
}
