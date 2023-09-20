package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var DB *sql.DB

func ConnectDB() {
	// Connect to the database
	db, err := sql.Open("postgres", "postgres://testuser:testpass@localhost/database?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)
	DB = db
	fmt.Println("Connected to database")
}
