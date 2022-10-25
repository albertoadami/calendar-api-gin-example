package repository

import "github.com/albertoadami/calendar-api-gin-example/pkg/repository/entity"

type UserRepository interface {
	CreateUser(user entity.UserEntity) (uint, error)
	ExistUserByEmail(email string) bool
}
