package controller

import (
	"github.com/albertoadami/calendar-api-gin-example/pkg/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)



type UserServiceMock struct {
	mock.Mock
}

func (c *UserServiceMock) CreateUser(req json.CreateUserRequest) (uint, error) {
	args := c.Called(req)
	return uint(args.Int(0)), args.Error(1)
}

func setUpRouter() *gin.Engine {
	r := gin.Default()

	userService := new(UserServiceMock)
	userService.On("CreateUser", 123).Return(true, nil)

	userController := NewUserController(userService)

	r.POST("users", userController.CreateUserHandler)
	return r
}
