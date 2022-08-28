package handler

import (
	"github.com/gin-gonic/gin"

	"hirotoohira.link/time-bucket/handler/api_handler"
)

func Route(r *gin.Engine) {
	r.GET("/health_check", api_handler.HandleHealthCheck)
}
