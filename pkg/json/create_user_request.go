package json

import "github.com/albertoadami/calendar-api-gin-example/pkg/domain"

type CreateUserRequest struct {
	FirstName string        `json:"firstname" binding:"required"`
	LastName  string        `json:"lastname" binding:"required"`
	Email     string        `json:"email" binding:"required"`
	Password  string        `json:"password" binding:"required"`
	Gender    domain.Gender `json:"gender" binding:"required"`
}
