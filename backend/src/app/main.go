package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"hirotoohira.link/time-bucket/infrastracture/repository_impl"

	"hirotoohira.link/time-bucket/handler"
)

func main() {
	fmt.Println("server start...")

	dbmap, err := repository_impl.InitDB()
	if err != nil {
		log.Fatalf("failed to init db: %v\n", err)
	}

	r := gin.Default()
	handler.Route(r, dbmap)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run gin server: %v\n", err)
	}
}
