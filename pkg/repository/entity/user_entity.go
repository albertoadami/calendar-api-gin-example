package entity

import (
	"time"

	"github.com/albertoadami/calendar-api-gin-example/pkg/domain"
)

type UserEntity struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	Gender    domain.Gender
	Enabled   bool
	CreatedAt time.Time
}
