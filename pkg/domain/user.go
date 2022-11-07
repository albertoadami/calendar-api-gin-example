package domain

import "time"

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Gender    Gender
	Enabled   bool
	CreatedAt time.Time
}
