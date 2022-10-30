package service

import (
	"github.com/albertoadami/calendar-api-gin-example/pkg/converter"
	customerrors "github.com/albertoadami/calendar-api-gin-example/pkg/errors"
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
		emailInUserError := &customerrors.EmailAlreadyInUseError{Email: req.Email}
		return 0, emailInUserError
	} else {
		userEntity := converter.FromCreateUserRequestToDomain(req)
		return c.userRepository.CreateUser(&userEntity)
	}

}
