package main

import "github.com/gin-gonic/gin"

func NamesHandler(c *gin.Context) {
	data := gin.H{
		"Title": "Vos noms de domaine",
	}

	ShowPage(c, "names", data)
}