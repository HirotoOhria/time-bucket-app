package repository

import (
	"hirotoohira.link/time-bucket/domain/entity"
	"hirotoohira.link/time-bucket/domain/value"
)

type BucketItemRepository interface {
	Get(id value.BucketItemID) (*entity.BucketItem, error)
	List() ([]*entity.BucketItem, error)
	Insert(bucketItem *entity.BucketItem) error
}
