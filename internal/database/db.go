package database

import (
	"log"
	"routes-service/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=" + config.Config("DBHost") + " user=" + config.Config("DBUser") + " password=" + config.Config("DBPassword") +
		" dbname=" + config.Config("DbName") + " port=" + string(config.Config("Port")) + " sslmode=" + config.Config("SSlMode")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	DB = db
}
