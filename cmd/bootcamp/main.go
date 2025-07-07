package main

import (
	"database/sql"
	"log"

	"github.com/melikadaryabak/bootcamp/internal/application/services"
	"github.com/melikadaryabak/bootcamp/internal/handler"
)

func main() {
	// Connect to MySQL
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/amacodecamp")

	// Check database connection
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	// Check database connection
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	srvc := services.NewServices(db)

	err = handler.NewServer("8080", srvc)
	if err != nil {
		log.Fatalf("error server: ", err)
	}
}
