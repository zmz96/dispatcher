package db

import (
	"fmt"
	"log"

	. "dispatcher/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db_url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		Config.Database.Host, Config.Database.User, Config.Database.Password,
		Config.Database.DBName, Config.Database.Port, Config.Database.SSLMode, Config.Database.TimeZone)
	db, err := gorm.Open(postgres.Open(db_url))
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}
	return db
}
