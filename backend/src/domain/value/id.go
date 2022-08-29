package value

type BucketItemID string

func NewBucketItemID(id string) (BucketItemID, error) {
	return BucketItemID(id), nil
}

func (id BucketItemID) Value() string {
	return string(id)
}
