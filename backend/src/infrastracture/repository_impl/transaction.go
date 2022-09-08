package repository_impl

import (
	"github.com/go-gorp/gorp"

	"hirotoohira.link/time-bucket/domain/repository"
)

type transactionImpl struct {
	tx *gorp.Transaction
}

func newTransaction(tx *gorp.Transaction) repository.Transaction {
	return &transactionImpl{
		tx: tx,
	}
}

func (t *transactionImpl) BucketItem() repository.BucketItemRepository {
	return NewBucketItemRepository(t.tx)
}
