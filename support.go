package main

import "github.com/gin-gonic/gin"

func SupportHandler(c *gin.Context) {
	data := gin.H{
		"Title": "Contacter le support",
	}

	ShowPage(c, "support", data)
}