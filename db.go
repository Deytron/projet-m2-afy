package main

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func CheckDb() {
	// Check if the database file exists
	var err error
		Db, err = sql.Open("sqlite3", "/app/db.sqlite")
		if err != nil {
			log.Fatal("Error connecting to database:", err)
		}

		if err != nil {
			log.Fatal("Error creating tables:", err)
		}

	// Initialize database connection
	var eerr error
	herr := CreateTables()
	Db, eerr = sql.Open("sqlite3", "db.sqlite")
	if eerr != nil {
		log.Fatal("Error connecting to database:", herr)
	}

	// Check if the connection to the database is successful
	perr := Db.Ping()
	if perr != nil {
		log.Fatal("Error pinging database:", perr)
	}
}
