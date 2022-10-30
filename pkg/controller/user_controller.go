package controller

import (
	"errors"
	"net/http"
	log "github.com/sirupsen/logrus"

	customerrors "github.com/albertoadami/calendar-api-gin-example/pkg/errors"
	"github.com/albertoadami/calendar-api-gin-example/pkg/json"
	"github.com/albertoadami/calendar-api-gin-example/pkg/service"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserRoutes {
	routes := &UserRoutes{userService: userService}

	return routes
}

func (r *UserRoutes) CreateUserHandler(c *gin.Context) {

	var req json.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, err := r.userService.CreateUser(req)
	if err != nil {
		var emailInUserError *customerrors.EmailAlreadyInUseError
		if errors.As(err, &emailInUserError) {
			c.Writer.WriteHeader(http.StatusConflict)
		} else {
			log.Errorf("Received an unexpected error: %s, returning 500 status code", err)
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		c.Writer.WriteHeader(http.StatusCreated)
	}

}
