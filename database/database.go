package database

import (
	"database/sql"
	"log"
	"os"
)

var DbConn *sql.DB

func SetupDatabase() {
	var err error

	DbConn, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
}
