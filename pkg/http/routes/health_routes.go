package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthRoutes(r *gin.Engine) {
	r.GET("/health", healthHandler)
}

func healthHandler(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNoContent)
}
