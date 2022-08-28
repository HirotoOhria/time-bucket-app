package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"hirotoohira.link/time-bucket/handler"
)

func main() {
	fmt.Println("hello world")

	r := gin.Default()
	handler.Route(r)

	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
