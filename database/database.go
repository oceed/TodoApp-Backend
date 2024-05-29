package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=1212 dbname=todoapp sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database Connected")
}
