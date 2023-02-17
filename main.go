package main

import (
	"fmt"
	"log"

	"github.com/clean-architecture/controller"
	"github.com/clean-architecture/entity"
	"github.com/clean-architecture/repository"
	"github.com/clean-architecture/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=reddit-api port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	fmt.Println("Connected to postgres successfully!")

	// Migrate the schema
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal("failed to migrate schema:", err)
	}
	fmt.Println("Database migrations occurred suucessfully!")

	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserService(*userRepository)
	userController := controller.NewUserController(*userService)

	r := gin.Default()

	r.POST("/users", userController.CreateUser)
	// r.GET("/users/:id", userController.GetUser)
	// r.GET("/users", userController.ListUsers)

	if err := r.Run(); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
