package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/signin")
	r.POST("/signup")
	r.GET("/signout")

	r.GET("/authorize")
	r.GET("/token")
	r.GET("/me")

	if err := r.Run(":8080"); err != nil {
		log.Println("Failed to start server:", err.Error())
	}
}
