package service

import (
	"testing"

	"github.com/albertoadami/calendar-api-gin-example/pkg/repository/entity"
)

type MockSuccessUserRepository struct{}

func (r *MockSuccessUserRepository) CreateUser(user *entity.UserEntity) (uint, error) {
	return 123, nil
}
func (r *MockSuccessUserRepository) ExistUserByEmail(email string) bool {
	return false
}

func TestCreateUserSuccess(t *testing.T) {
	repo := &MockSuccessUserRepository{}
	userService := NewUserService(repo)

}
