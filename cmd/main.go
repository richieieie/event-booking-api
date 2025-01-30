package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/richieieie/event-booking/cmd/app"
	database "github.com/richieieie/event-booking/internal/db"
)

func main() {
	err := godotenv.Load("/Users/trung/Workspace/go/api-with-gin/config/env/auth.env")
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	err = database.InitDb()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	defer func() {
		sqlDB, err := database.Db.DB()
		if err != nil {
			log.Printf("Failed to get underlying database connection: %v", err)
			return
		}
		sqlDB.Close()
	}()

	err = app.Run()
	if err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
