package converter

import (
	"testing"

	"github.com/albertoadami/calendar-api-gin-example/pkg/domain"
	"github.com/albertoadami/calendar-api-gin-example/pkg/json"
	"github.com/stretchr/testify/assert"
)

func TestFromCreateUserRequestToDomain(t *testing.T) {

	var createUserRequest = json.CreateUserRequest{FirstName: "name", LastName: "last", Email: "name@last.com", Password: "password", Gender: domain.Male}

	result := FromCreateUserRequestToDomain(createUserRequest)

	assert.Equal(t, createUserRequest.FirstName, result.FirstName)
	assert.Equal(t, createUserRequest.LastName, result.LastName)
	assert.Equal(t, createUserRequest.Email, result.Email)
	assert.Equal(t, createUserRequest.Password, result.Password)
	assert.Equal(t, createUserRequest.Gender, result.Gender)
	assert.Equal(t, false, result.Enabled)

}
