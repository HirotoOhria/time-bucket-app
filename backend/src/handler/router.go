package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	"hirotoohira.link/time-bucket/handler/api_handler"

	"hirotoohira.link/time-bucket/injector"
)

func Route(r *gin.Engine, dbmap *gorp.DbMap) {
	r.GET("/health_check", api_handler.HandleHealthCheck)

	h := injector.InjectBucketItemHandler(dbmap)
	r.GET("/bucket_item/:id", h.Get)
	r.POST("/bucket_item", h.Create)
}
