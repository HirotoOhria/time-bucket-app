package injector

import (
	"github.com/go-gorp/gorp"
	"hirotoohira.link/time-bucket/domain/repository"
	"hirotoohira.link/time-bucket/infrastracture/repository_impl"
)

func InjectBucketItemRepository(dbmap *gorp.DbMap) repository.BucketItemRepository {
	return repository_impl.NewBucketItemRepositoryImpl(dbmap)
}
