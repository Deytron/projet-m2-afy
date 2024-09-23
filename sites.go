package main

import "github.com/gin-gonic/gin"

func SitesHandler(c *gin.Context) {
	data := gin.H{
		"Title": "Vos sites hébergés",
	}

	ShowPage(c, "sites", data)
}