package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func SetupDatabase() {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")

	var err error
	var config gorm.Config
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	if os.Getenv("ENABLE_GORM_LOGGER") != "" {
		config = gorm.Config{}
	} else {
		config = gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}

	DB, err = gorm.Open(postgres.Open(dsn), &config)

	if err != nil {
		log.Fatal(err)
		panic("Failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
}
