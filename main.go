package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	r := gin.Default()

	// Check if the database file exists
	if _, err := os.Stat("./db.sqlite"); os.IsNotExist(err) {
		fmt.Println("Database does not exist. Creating tables...")
		var err error
		db, err = sql.Open("sqlite3", "db.sqlite")
		if err != nil {
			log.Fatal("Error connecting to database:", err)
		}
		defer db.Close()

		err = CreateTables(db)
		if err != nil {
			log.Fatal("Error creating tables:", err)
		}
		fmt.Println("Database and tables created successfully.")
	} else {
		fmt.Println("Database already exists.")
	}

	// Initialize database connection
	var err error
	db, err = sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	r.StaticFile("/logo-light.png", "./assets/logo-light.png")
	r.StaticFile("/style.css", "./style/style.css")

	r.HTMLRender = ginview.Default()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "AFY",
		})
	})

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup", gin.H{
			"title": "Inscription",
		})
	})

	r.POST("/login", loginHandler)

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login", gin.H{
			"title": "Connexion",
		})
	})

	r.POST("/signup", registerHandler)

	r.Run(":8081")
}
