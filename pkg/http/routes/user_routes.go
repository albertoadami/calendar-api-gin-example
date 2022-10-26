package routes

import (
	"net/http"

	"github.com/albertoadami/calendar-api-gin-example/pkg/http/json"
	"github.com/albertoadami/calendar-api-gin-example/pkg/service"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userService *service.UserService
	c           *gin.Context
}

func NewUserRoutes(userService *service.UserService, c *gin.Context) *UserRoutes {
	routes := &UserRoutes{userService: userService, c: c}

	return routes
}

func InitRoutes(r *gin.Engine) {
	r.POST("/users", createUserHandler)
}

func createUserHandler(c *gin.Context) {

	var req json.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

}
