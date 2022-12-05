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

	return &a, nil
}

func (db accountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := db.client.Begin()
	if err != nil {
		logger.Error("error starting db transaction for account transaction " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	amount := t.Amount

	if t.Type == Withdrawal {
		amount *= -1
	}

	updateBalanceSql := "update accounts set amount = amount + ? where account_id = ?"

	res, err := tx.Exec(updateBalanceSql, amount, t.AccountId)
	if err != nil {
		tx.Rollback()
		logger.Error("error updating balance", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	rows, err := res.RowsAffected()
	if err != nil || rows != 1 {
		tx.Rollback()
		logger.Error("unexpected rows affected", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	saveSql := "insert into transactions (account_id, amount, transaction_date, transaction_type) values (?, ?, ?, ?)"

	tRes, err := tx.Exec(saveSql, t.AccountId, t.Amount, t.Time, t.Type)
	if err != nil {
		tx.Rollback()
		logger.Error("error completing transaction", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	id, err := tRes.LastInsertId()
	if err != nil {
		tx.Rollback()
		logger.Error("error getting last inserted transaction", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	tx.Commit()

	t.Id = strconv.FormatInt(id, 10)

	return &t, nil
}

func (d accountRepositoryDb) GetAccount(accountId string, customerId string) (*Account, *errs.AppError) {
	getSql := "select account_id, customer_id, amount, status from accounts where account_id = ? and customer_id = ?"

	var a Account
	err := d.client.Get(&a, getSql, accountId, customerId)

	if err == sql.ErrNoRows {
		return nil, errs.NewNotFoundError("account not found for customer")
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
