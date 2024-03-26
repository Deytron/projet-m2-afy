package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// User represents a user in the database
type User struct {
	ID       int
	Username string
	Password string
}

var db *sql.DB

func CreateTables(db *sql.DB) error {
	// Create users table
	_, err := db.Exec(`CREATE TABLE users (
		ID INTEGER PRIMARY KEY,
		username TEXT,
		password TEXT
	)`)
	if err != nil {
		return err
	}
	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Serve login/register form
	tpl := template.Must(template.ParseFiles("index.html"))
	tpl.Execute(w, nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Insert user into the database
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Redirect to login page after successful registration
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Query the database for the user
	row := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
	user := User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		log.Println(err)
		return
	}

	// Check if the password matches
	if password != user.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Successful login
	fmt.Fprintf(w, "Welcome, %s!", user.Username)
}
