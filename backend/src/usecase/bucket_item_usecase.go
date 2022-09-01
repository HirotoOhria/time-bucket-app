package usecase

import (
	"hirotoohira.link/time-bucket/domain/entity"
	"hirotoohira.link/time-bucket/domain/repository"
	"hirotoohira.link/time-bucket/domain/value"
	"hirotoohira.link/time-bucket/handler/api_io"
	"hirotoohira.link/time-bucket/util"
)

type BucketItemUsecase struct {
	bucketItemRepository repository.BucketItemRepository
}

func NewBucketItemUsecase(bucketItemRepository repository.BucketItemRepository) *BucketItemUsecase {
	return &BucketItemUsecase{
		bucketItemRepository: bucketItemRepository,
	}
}

func (u *BucketItemUsecase) Get(input *api_io.BucketItemGetInput) (*api_io.BucketItemGetOutput, error) {
	buketItemID, err := value.NewBucketItemID(input.ID)
	if err != nil {
		return nil, err
	}

	bucketItem, err := u.bucketItemRepository.Get(buketItemID)
	if err != nil {
		return nil, err
	}

	return u.convertBucketItemToOutput(bucketItem), nil
}

func (u *BucketItemUsecase) List() (api_io.BucketItemListOutput, error) {
	bis, err := u.bucketItemRepository.List()
	if err != nil {
		return nil, err
	}

	res := make(api_io.BucketItemListOutput, len(bis))
	for i, bi := range bis {
		res[i] = &api_io.BucketItemGetOutput{
			ID:   bi.ID.Value(),
			Name: bi.Name,
		}
	}

	return res, nil
}

func (u *BucketItemUsecase) Create(input *api_io.BucketItemCreateInput) error {
	bucketItemID, err := value.NewBucketItemID(util.NewULID())
	if err != nil {
		return err
	}

	bucketItem := &entity.BucketItem{
		ID:   bucketItemID,
		Name: input.Name,
	}

	return u.bucketItemRepository.Insert(bucketItem)
}

func (u *BucketItemUsecase) convertBucketItemToOutput(bucketItem *entity.BucketItem) *api_io.BucketItemGetOutput {
	return &api_io.BucketItemGetOutput{
		ID:   bucketItem.ID.Value(),
		Name: bucketItem.Name,
	}
}
