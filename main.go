package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Check for SQL File
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the database file exists
	if _, err := os.Stat("./db.sqlite"); os.IsNotExist(err) {
		err := CreateTables(db)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Database and tables created successfully.")
	} else {
		fmt.Println("Database already exists.")
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

	r.Run(":8081")
}
