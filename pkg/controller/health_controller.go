package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealthHandler(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNoContent)
}
