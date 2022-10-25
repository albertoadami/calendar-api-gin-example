package main

import (
	"github.com/albertoadami/calendar-api-gin-example/pkg/http/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/config/.env")
	viper.ReadInConfig()

	router := gin.Default()

	routes.HealthRoutes(router)

	router.Run()

}
