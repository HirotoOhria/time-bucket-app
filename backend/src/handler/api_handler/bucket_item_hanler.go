package api_handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"hirotoohira.link/time-bucket/handler/api_io"
	"hirotoohira.link/time-bucket/usecase"
)

type BucketItemHandler struct {
	bucketItemUsecase *usecase.BucketItemUsecase
}

func NewBucketItemHandler(bucketItemUsecase *usecase.BucketItemUsecase) *BucketItemHandler {
	return &BucketItemHandler{
		bucketItemUsecase: bucketItemUsecase,
	}
}

func (h *BucketItemHandler) Get(c *gin.Context) {
	i := h.getGetInput(c)

	o, err := h.bucketItemUsecase.Get(i)
	if err != nil {
		fmt.Printf("failed to get bucket item: %v\n", err)
		responseInternal(c, err)
		return
	}

	responseOK(c, o)
}

func (h *BucketItemHandler) getGetInput(c *gin.Context) *api_io.BucketItemGetInput {
	return &api_io.BucketItemGetInput{
		ID: c.Param("id"),
	}
}

func (h *BucketItemHandler) Create(c *gin.Context) {
	i := h.getCreateInput(c)

	if err := h.bucketItemUsecase.Create(i); err != nil {
		fmt.Printf("failed to create bucket item: %v\n", err)
		responseInternal(c, err)
		return
	}

	responseSuccess(c)
}

func (h *BucketItemHandler) getCreateInput(c *gin.Context) *api_io.BucketItemCreateInput {
	return &api_io.BucketItemCreateInput{
		Name: c.PostForm("name"),
	}
}
