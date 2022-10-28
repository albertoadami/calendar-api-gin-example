package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthController(r *gin.Engine) {
	r.GET("/health", healthHandler)
}

func healthHandler(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNoContent)
}
