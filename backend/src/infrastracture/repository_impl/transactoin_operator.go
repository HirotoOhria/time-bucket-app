package repository_impl

import (
	"github.com/go-gorp/gorp"

	"hirotoohira.link/time-bucket/domain/repository"
)

type transactionOperatorImpl struct {
	db *gorp.DbMap
}

func NewTransactionOperator(db *gorp.DbMap) repository.TransactionOperator {
	return &transactionOperatorImpl{
		db: db,
	}
}

func (t *transactionOperatorImpl) RunTransaction(txFunc repository.TxFunc) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	if err := txFunc(newTransaction(tx)); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
