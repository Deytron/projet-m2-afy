package main

import "github.com/gin-gonic/gin"

var homego = "home.go"

func HomeHandler(c *gin.Context) {
	data := gin.H{
		"Title": "Accueil",
	}

	ShowPage(c, "home", data)
}

