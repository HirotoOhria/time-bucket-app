package repository

type TxFunc func(tx Transaction) error

type TransactionOperator interface {
	RunTransaction(tx TxFunc) error
}

type Transaction interface {
	BucketItem() BucketItemRepository
}
