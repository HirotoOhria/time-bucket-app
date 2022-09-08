package repository_impl

import (
	"fmt"

	"github.com/go-gorp/gorp"

	"hirotoohira.link/time-bucket/domain/entity"
	"hirotoohira.link/time-bucket/domain/repository"
	"hirotoohira.link/time-bucket/domain/value"
	"hirotoohira.link/time-bucket/infrastracture/model"
)

type bucketItemRepositoryImpl struct {
	db gorp.SqlExecutor
}

func NewBucketItemRepository(db gorp.SqlExecutor) repository.BucketItemRepository {
	return &bucketItemRepositoryImpl{
		db: db,
	}
}

func (r *bucketItemRepositoryImpl) Get(id value.BucketItemID) (*entity.BucketItem, error) {
	var bucketItem *model.BucketItem

	query := fmt.Sprintf(`
SELECT %s
FROM bucket_items
WHERE id = ?
`, bucketItemColumns)

	if err := r.db.SelectOne(&bucketItem, query, id); err != nil {
		return nil, wrapSQLError(err)
	}

	return convertBucketItemToEntity(bucketItem)
}

func (r *bucketItemRepositoryImpl) List() ([]*entity.BucketItem, error) {
	var bucketItems []*model.BucketItem

	query := fmt.Sprintf(`
SELECT %s
FROM bucket_items
`, bucketItemColumns)

	if _, err := r.db.Select(&bucketItems, query); err != nil {
		return nil, wrapSQLError(err)
	}

	return convertBucketItemsToEntity(bucketItems)
}

func (r *bucketItemRepositoryImpl) Insert(bucketItem *entity.BucketItem) error {
	m := convertBucketITemToModel(bucketItem)

	if err := r.db.Insert(m); err != nil {
		return wrapSQLError(err)
	}

	return nil
}

func convertBucketItemToEntity(bucketItem *model.BucketItem) (*entity.BucketItem, error) {
	id, err := value.NewBucketItemID(bucketItem.ID)
	if err != nil {
		return nil, err
	}

	return &entity.BucketItem{
		ID:   id,
		Name: bucketItem.Name,
	}, nil
}

func convertBucketItemsToEntity(bucketItems []*model.BucketItem) ([]*entity.BucketItem, error) {
	var err error

	res := make([]*entity.BucketItem, len(bucketItems))
	for i, bi := range bucketItems {
		res[i], err = convertBucketItemToEntity(bi)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func convertBucketITemToModel(bucketItem *entity.BucketItem) *model.BucketItem {
	return &model.BucketItem{
		ID:   bucketItem.ID.Value(),
		Name: bucketItem.Name,
	}
}
