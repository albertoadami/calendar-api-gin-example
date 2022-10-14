package repository

type UserRepository interface {
	CreateUser(s string) int
}
