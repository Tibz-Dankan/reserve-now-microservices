package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Db() *sql.DB {
	connectionString := os.Getenv("DATABASE_URL_RESERVE_NOW")
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal("error occurred while connecting to the database", err)
	}

	fmt.Println("Connected to database successfully")
	return db
}
