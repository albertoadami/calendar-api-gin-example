package converter

import (
	"time"

	"github.com/albertoadami/calendar-api-gin-example/pkg/json"
	"github.com/albertoadami/calendar-api-gin-example/pkg/repository/entity"
)

func FromCreateUserRequestToDomain(req json.CreateUserRequest) entity.UserEntity {
	user := entity.UserEntity{FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Password: req.Password, Gender: req.Gender, Enabled: false, CreatedAt: time.Now()}
	return user
}
