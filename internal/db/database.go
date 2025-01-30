package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/richieieie/event-booking/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// For using database/sql with postgres driver, reference to this link: https://pkg.go.dev/github.com/lib/pq

var Db *gorm.DB

func InitDb() error {
	if Db == nil {
		err := godotenv.Load("/Users/trung/Workspace/go/api-with-gin/config/env/postgres.env")
		if err != nil {
			return fmt.Errorf("failed to load environment variables: %w", err)
		}
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)

		Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			QueryFields: true,
		})
		if err != nil {
			return fmt.Errorf("failed to connect to database: %w", err)
		}

		sqlDB, err := Db.DB()
		if err != nil {
			return fmt.Errorf("failed to get underlying database connection: %w", err)
		}
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)

		if err := Db.AutoMigrate(&model.Event{}, &model.User{}); err != nil {
			return fmt.Errorf("failed to migrate database: %w", err)
		}

		log.Println("Database initialized successfully")
	}
	return nil
}
