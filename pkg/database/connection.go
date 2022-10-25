package database

import (
	"fmt"

	"github.com/albertoadami/calendar-api-gin-example/pkg/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func init() {
	config, error := config.LoadConfig()

	if error != nil {
		panic("error during loading of configuration, impossibile to init database connection")
	} else {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.DatabaseHost, config.DatabaseUser, config.DatabasePassword, config.DatabaseName, config.DatabasePort)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(fmt.Sprintf("Error connecting to database : error=%v", err))
		}

		log.Info("Connected to database ", config.DatabaseName)
		database = db
	}

}

func GetConnection() *gorm.DB {
	return database
}
