package api_handler

import (
	"github.com/gin-gonic/gin"
)

func HandleHealthCheck(c *gin.Context) {
	responseOK(c, "ok")
}
