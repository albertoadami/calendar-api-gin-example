package routes

import (
	"github.com/albertoadami/calendar-api-gin-example/pkg/service"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userService service.UserService
}

func InitRoutes(r *gin.Engine) {
	r.POST("/users")
}
