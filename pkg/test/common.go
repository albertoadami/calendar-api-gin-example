package test

import (
	"github.com/albertoadami/calendar-api-gin-example/pkg/domain"
	"github.com/albertoadami/calendar-api-gin-example/pkg/json"
)

func GetCreateUserRequest() *json.CreateUserRequest {
	return &(json.CreateUserRequest{FirstName: "test", LastName: "test", Email: "test@test.it", Password: "password", Gender: domain.Male})
}
