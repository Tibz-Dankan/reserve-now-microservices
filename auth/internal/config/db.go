package config

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db() *gorm.DB {

	// dsn := os.Getenv("DSN_RESERVE_NOW")
	dsn := "host=localhost user=postgres password=ourpassword dbname=reserve_now_microservice_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	fmt.Println("Connected to database successfully")
	return db
}

// gorm migrate create -name=create-users
