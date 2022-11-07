package repository

import (
	"testing"
	"time"

	"github.com/albertoadami/calendar-api-gin-example/pkg/domain"
	"github.com/albertoadami/calendar-api-gin-example/pkg/repository/entity"
	"github.com/albertoadami/calendar-api-gin-example/pkg/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testPostgres = test.NewTestDatabase()
var db, _ = gorm.Open(postgres.Open(testPostgres.Dsn()))
var userRepository = NewUserRepository(db)

var testEmail = "test@test.it"
var newUserEntity = entity.UserEntity{FirstName: "test", LastName: "test", Email: testEmail, Password: "test", Gender: domain.Male, Enabled: false, CreatedAt: time.Now()}

func TestCreateUser(t *testing.T) {
	_, err := userRepository.CreateUser(&newUserEntity) //create first unique user
	assert.Equal(t, nil, err)
}
