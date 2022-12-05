package transaction

import (
	"strconv"

	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type transactionRepositoryDb struct {
	client *sqlx.DB
}

func (db transactionRepositoryDb) Save(t Transaction) (*Transaction, *errs.AppError) {
	saveSql := "insert into transactions (account_id, amount, transaction_date, transaction_type) values (?, ?, ?, ?)"

	res, err := db.client.Exec(saveSql, t.AccountId, t.Amount, t.Time, t.Type)
	if err != nil {
		logger.Error("error completing transaction", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Error("error getting last inserted transaction", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	t.Id = strconv.FormatInt(id, 10)

	return &t, nil
}

func NewDbRepository(db *sqlx.DB) *transactionRepositoryDb {
	return &transactionRepositoryDb{
		client: db,
	}
}
