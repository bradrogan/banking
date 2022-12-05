package account

import (
	"database/sql"
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
	saveSql := "insert into accounts (account_type, amount, customer_id, opening_date, status) values (?, ?, ?, ?, ?)"

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

func (d accountRepositoryDb) UpdateBalance(a Account) (*Account, *errs.AppError) {
	updateBalanceSql := "update accounts set amount = ? where account_id = ?"

	res, err := d.client.Exec(updateBalanceSql, a.Amount, a.Id)
	if err != nil {
		logger.Error("error updating balance", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	rows, err := res.RowsAffected()
	if err != nil || rows != 1 {
		logger.Error("error getting rows affected", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return &a, nil
}

func (d accountRepositoryDb) GetAccount(id string) (*Account, *errs.AppError) {
	getSql := "select account_id, customer_id, amount, status from accounts where account_id = ?"

	var a Account
	err := d.client.Get(&a, getSql, id)

	if err == sql.ErrNoRows {
		return nil, errs.NewNotFoundError("account not found")
	}
	if err != nil {
		logger.Error("error while scanning account table", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database errror" + err.Error())
	}
	return &a, nil
}

func NewDbRepository(db *sqlx.DB) accountRepositoryDb {
	return accountRepositoryDb{client: db}
}
