package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hiepphatle1104/nextmesh/internal/handler"
	"github.com/hiepphatle1104/nextmesh/internal/middleware"
	"github.com/hiepphatle1104/nextmesh/internal/service"
	"github.com/hiepphatle1104/nextmesh/internal/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=admin password=nextmesh dbname=nextmesh_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
	}))

	store := storage.NewStorage(db)
	service := service.NewService(store)
	controller := handler.NewController(service)

	r.POST("/signin", controller.LoginHandler)
	r.POST("/signup", controller.RegisterHandler)

	r.Use(middleware.SessionMiddleware)
	r.GET("/signout", controller.LogoutHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err.Error())
	}
}
