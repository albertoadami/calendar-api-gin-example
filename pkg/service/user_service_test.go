package service

import (
	"fmt"
	"testing"

	errors "github.com/albertoadami/calendar-api-gin-example/pkg/errors"
	"github.com/albertoadami/calendar-api-gin-example/pkg/repository/entity"
	"github.com/albertoadami/calendar-api-gin-example/pkg/test"
	"github.com/stretchr/testify/assert"
)

var userId uint = 123
var createUserRequest = test.GetCreateUserRequest()

type MockSuccessUserRepository struct{}

func (r *MockSuccessUserRepository) CreateUser(user *entity.UserEntity) (uint, error) {
	return userId, nil
}
func (r *MockSuccessUserRepository) ExistUserByEmail(email string) bool {
	return false
}

type MockFailureUserRepository struct{ exist bool }

func (r *MockFailureUserRepository) CreateUser(user *entity.UserEntity) (uint, error) {
	return 0, fmt.Errorf("Some error occurred")
}
func (r *MockFailureUserRepository) ExistUserByEmail(email string) bool {
	return r.exist
}

func TestCreateUserSuccess(t *testing.T) {

	repo := &MockSuccessUserRepository{}
	userService := NewUserService(repo)

	result, err := userService.CreateUser(createUserRequest)
	assert.Equal(t, userId, result)
	assert.Equal(t, nil, err)
}

func TestCreateUserAlreadyInUse(t *testing.T) {
	var noUserId uint = 0

	repo := &MockFailureUserRepository{true}
	userService := NewUserService(repo)

	result, err := userService.CreateUser(createUserRequest)
	assert.Equal(t, noUserId, result)
	assert.Equal(t, &errors.EmailAlreadyInUseError{Email: createUserRequest.Email}, err)
}

func TestCreateUserError(t *testing.T) {
	var noUserId uint = 0

	repo := &MockFailureUserRepository{false}
	userService := NewUserService(repo)

	result, err := userService.CreateUser(createUserRequest)
	assert.Equal(t, noUserId, result)
	assert.Equal(t, fmt.Errorf("Some error occurred"), err)

}
