package injector

import (
	"github.com/go-gorp/gorp"
	"hirotoohira.link/time-bucket/usecase"
)

func InjectBucketItemUsecase(dbmap *gorp.DbMap) *usecase.BucketItemUsecase {
	return usecase.NewBucketItemUsecase(InjectBucketItemRepository(dbmap))
}
