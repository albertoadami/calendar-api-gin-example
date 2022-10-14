package service

import "github.com/albertoadami/calendar-api-gin-example/pkg/repository"

type UserService struct {
	userRepository repository.UserRepository
}
