package errors

import "fmt"

type EmailAlreadyInUseError struct {
	Email string
}

func (err *EmailAlreadyInUseError) Error() string {
	return fmt.Sprintf("User with email %s already exists", err.Email)
}