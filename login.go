package main

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var logingo = "login.go"

func CreateTables() error {
	// Create users table
	_, err := Db.Exec(`CREATE TABLE IF NOT EXISTS users (
		ID TEXT PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	)`)
	_, berr := Db.Exec(`CREATE TABLE IF NOT EXISTS vms (
		ID TEXT PRIMARY KEY,
		vmname TEXT,
		ipaddress TEXT NULL,
		description TEXT NULL,
		sku TEXT,
		datacenter TEXT,
		username TEXT
	)`)
	_, cerr := Db.Exec(`CREATE TABLE IF NOT EXISTS site (
		ID TEXT PRIMARY KEY,
		sitename TEXT NOT NULL UNIQUE,
		userId TEXT UNIQUE,

		CONSTRAINT fk_userId
      		FOREIGN KEY(userId) 
        		REFERENCES users(ID)
	)`)
	_, derr := Db.Exec(`CREATE TABLE IF NOT EXISTS dns (
		ID TEXT PRIMARY KEY,
		domainName TEXT,
		extension TEXT,
		userId TEXT UNIQUE,

		CONSTRAINT fk_userId
      		FOREIGN KEY(userId) 
        		REFERENCES users(ID)
	)`)
	if err != nil {
		return err
	}
	if berr != nil {
		return err
	}
	if cerr != nil {
		return err
	}
	if derr != nil {
		return err
	}

	return nil
}

func RegisterHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		email := c.PostForm("email")
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Check if user with the given username or eemail already exists
		var us User
		err := Db.QueryRow("SELECT id, email, username FROM users WHERE email = ? OR username = ?", email, username).Scan(&us.ID, &us.Email, &us.Username)
		if err != nil && err != sql.ErrNoRows {
			NonFatal(err, logingo, "Querying register error")
		}

		if us.ID != "" {
			ErrorMessage = "Erreur : l'utilisateur existe déjà"
		}

		// Generate UUID for the user
		userID := uuid.New().String()

		// Hash password
		bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
		hashedPassword := string(bytes)

		// Insert user into the database
		_, err = Db.Exec("INSERT INTO users (id, email, username, password) VALUES (?, ?, ?, ?)", userID, email, username, hashedPassword)
		if err != nil {
			NonFatal(err, logingo, "Inserting new user in db.sqlite")
		}
		c.Redirect(301, "/login")
	}

	data := gin.H{
		"Title": "Créer un compte",
	}

	ShowPage(c, "register", data)
}

func LoginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		email := c.PostForm("email")
		password := c.PostForm("password")

		// Retrieve user details from the database based on the provided email
		var us User
		err := Db.QueryRow("SELECT id, email, username, password FROM users WHERE email = ?", email).Scan(&us.ID, &us.Email, &us.Username, &us.Password)
		if err != nil {
			if err == sql.ErrNoRows {
				ErrorMessage = "Erreur : utilisateur ou mot de passe incorrect"
			}
			NonFatal(err, logingo, "Querying login error")
		}

		// Compare hashed password with the one provided
		err = bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(password))
		if err != nil {
			ErrorMessage = "Erreur : utilisateur ou mot de passe incorrect"
		}

		SetSessionCookie(c, us.Username)
		c.Redirect(301, "/")
	}

	data := gin.H{
		"Title": "Connexion",
	}

	ShowPage(c, "login", data)
}
