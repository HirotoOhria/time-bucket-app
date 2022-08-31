package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"hirotoohira.link/time-bucket/handler"
	"hirotoohira.link/time-bucket/infrastracture/repository_impl"
)

func main() {
	fmt.Println("server start...")

	dbmap, err := repository_impl.InitDB()
	if err != nil {
		log.Fatalf("failed to init db: %v\n", err)
	}

	r := gin.Default()
	allowAllCORS(r)

	handler.Route(r, dbmap)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run gin server: %v\n", err)
	}
}

func allowAllCORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
}
