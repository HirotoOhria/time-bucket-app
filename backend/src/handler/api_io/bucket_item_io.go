package api_io

type BucketItemGetInput struct {
	ID string
}

type BucketItemGetOutput struct {
	ID   string
	Name string
}

type BucketItemCreateInput struct {
	Name string
}
