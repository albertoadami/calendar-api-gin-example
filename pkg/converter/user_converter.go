package converter

import (
	"time"

	"github.com/albertoadami/calendar-api-gin-example/pkg/domain"
	"github.com/albertoadami/calendar-api-gin-example/pkg/http/json"
)

func FromCreateUserRequestToUser(req json.CreateUserRequest) domain.User {
	user := domain.User{FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Password: req.Password, Gender: req.Gender, Enabled: false, CreatedAt: time.Now()}
	return user
}
