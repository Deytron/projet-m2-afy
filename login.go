package main

import (
	"database/sql"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

// User represents a user in the database
type User struct {
	ID       string
	Username string
	Password string
	email    string
}

const saltSize = 16

func CreateTables(db *sql.DB) error {
	// Create users table
	_, err := db.Exec(`CREATE TABLE users (
		ID TEXT PRIMARY KEY,
		username TEXT,
		password TEXT,
		email TEXT
	)`)
	if err != nil {
		return err
	}
	return nil
}

func registerHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Check if user with the given username or eemail already exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ? OR email = ?", username, email).Scan(&count)
	if err != nil {
		http.Error(w, "Error checking user existence", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if count > 0 {
		http.Error(w, "Username or email already exists", http.StatusBadRequest)
		return
	}

	// Generate UUID for the user
	userID := uuid.New().String()

	// Hash password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	hashedPassword := string(bytes)

	// Insert user into the database
	_, err = db.Exec("INSERT INTO users (id, email, username, password) VALUES (?, ?, ?, ?)", userID, email, username, hashedPassword)
	if err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Redirect to login page after successful registration
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func loginHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Query the database for the user
	row := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
	user := User{}
	var hashedPassword string
	err := row.Scan(&user.ID, &user.Username, &hashedPassword)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		log.Println(err)
		return
	}

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		log.Println(err)
		return
	}

	// Successful login
	// Set a cookie with the user's ID
	cookie := http.Cookie{
		Name:     "userID",
		Value:    user.ID,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	// Redirect to login page after successful login
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func checkAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the userID cookie exists
		cookie, err := r.Cookie("userID")
		if err != nil || cookie.Value == "" {
			// If the cookie doesn't exist, redirect the user to the login page
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// If the cookie exists, call the next handler
		next.ServeHTTP(w, r)
	})
}
