package service

import (
	"fmt"

	"github.com/albertoadami/calendar-api-gin-example/pkg/converter"
	"github.com/albertoadami/calendar-api-gin-example/pkg/json"
	"github.com/albertoadami/calendar-api-gin-example/pkg/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (c *UserService) CreateUser(req json.CreateUserRequest) (uint, error) {
	exists := c.userRepository.ExistUserByEmail(req.Email)

	if exists {
		return 0, fmt.Errorf("User with email %s already exists", req.Email)
	} else {
		userEntity := converter.FromCreateUserRequestToDomain(req)
		return c.userRepository.CreateUser(&userEntity)
	}

}
