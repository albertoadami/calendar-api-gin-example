package repository

import (
	"github.com/albertoadami/calendar-api-gin-example/pkg/repository/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.UserEntity) (uint, error)
	ExistUserByEmail(email string) bool
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db}
}

func (r *GormUserRepository) CreateUser(user *entity.UserEntity) (uint, error) {
	if err := r.db.Create(user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *GormUserRepository) ExistUserByEmail(email string) bool {
	var user entity.UserEntity
	return r.db.Where("email = ?", email).Find(&user).RowsAffected > 0
}
