package injector

import (
	"github.com/go-gorp/gorp"
	"hirotoohira.link/time-bucket/handler/api_handler"
)

func InjectBucketItemHandler(dbmap *gorp.DbMap) *api_handler.BucketItemHandler {
	return api_handler.NewBucketItemHandler(InjectBucketItemUsecase(dbmap))
}
