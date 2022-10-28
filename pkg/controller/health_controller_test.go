package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setUpRouter() *gin.Engine {
	r := gin.Default()
	HealthController(r)
	return r
}

func TestHealthController(t *testing.T) {
	router := setUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)

}
