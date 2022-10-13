package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/config/.env")
	viper.ReadInConfig()

	servicePort := viper.Get("SERVICE_PORT").(string)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusNoContent)
	})

	r.Run(fmt.Sprintf(":%s", servicePort))

}
