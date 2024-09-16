package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var Db *sql.DB

func CheckDb() {
	// Check if the database file exists
	if _, err := os.Stat("./db.sqlite"); os.IsNotExist(err) {
		fmt.Println("Database does not exist. Creating tables...")
		var err error
		Db, err = sql.Open("sqlite3", "Db.sqlite")
		if err != nil {
			log.Fatal("Error connecting to database:", err)
		}
		defer Db.Close()

		err = CreateTables(Db)
		if err != nil {
			log.Fatal("Error creating tables:", err)
		}
		fmt.Println("Database and tables created successfully.")
	} else {
		fmt.Println("Database already exists.")
	}

	// Initialize database connection
	var err error
	Db, err = sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer Db.Close()

	// Check if the connection to the database is successful
	err = Db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}
}
