package main

import (
	"github.com/albertoadami/calendar-api-gin-example/pkg/database"
	"github.com/albertoadami/calendar-api-gin-example/pkg/http/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	_ = database.GetConnection()

	router := gin.Default()

	routes.HealthRoutes(router)

	router.Run()

}
