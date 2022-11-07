package main

import (
	"github.com/albertoadami/calendar-api-gin-example/pkg/controller"
	"github.com/albertoadami/calendar-api-gin-example/pkg/database"
	"github.com/albertoadami/calendar-api-gin-example/pkg/repository"
	"github.com/albertoadami/calendar-api-gin-example/pkg/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Initialize Database connection...")
	connection := database.GetConnection()
	log.Info("Database connection completed successfully.")

	log.Info("Migrating Database structure...")
	database.MigrateDatabase()
	log.Info("Database structure completed successfully.")

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//repositories
	userRepository := repository.NewUserRepository(connection)

	//services
	userService := service.NewUserService(userRepository)

	router.GET("/health", controller.GetHealthHandler)
	//controllers
	userController := controller.NewUserController(userService)

	v1 := router.Group("/api/v1")
	v1.POST("/users", userController.CreateUserHandler)

	router.Run()

}
