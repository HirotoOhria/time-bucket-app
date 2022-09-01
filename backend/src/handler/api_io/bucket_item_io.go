package api_io

type BucketItemGetInput struct {
	ID string `json:"id"`
}

type BucketItemGetOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BucketItemListOutput []*BucketItemGetOutput

type BucketItemCreateInput struct {
	Name string `json:"name"`
}
