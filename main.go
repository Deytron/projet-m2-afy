package main

import (
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// ENABLE DEBUG HERE by setting it to true
var Debug = false

// Variables
var SuccessMessage string
var ErrorMessage string

var port = ":8081"
var maingo = "main.go"

func main() {
	// Load dotenv
	err := godotenv.Load()
	Fatal(err, maingo, "Loading .env")

	CheckDb()
	// Initialisation of router
	InitConfig()

	// Now you can use config.GetConfig().Router and config.GetConfig().Templates
	r := GetConfig().Router
	// r.Use(CORSMiddleware())
	SetRoutes(r)
	r.Use(CORSMiddleware())
	// Set limit for file upload
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Static files
	r.Static("/assets", "assets/")

	// Don't trus all proxies
	perr := r.SetTrustedProxies(nil)
	NonFatal(perr, maingo, "Setting trusted proxies")

	r.Run(port)
}
