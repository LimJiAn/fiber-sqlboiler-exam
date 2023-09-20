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
	var connStr string = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "testuser", "testpass", "testdb")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)
	DB = db
	fmt.Println("🙆‍♂️ Connected to database")
}
