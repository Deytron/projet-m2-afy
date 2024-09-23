package main

import "github.com/gin-gonic/gin"

func PolicyHandler(c *gin.Context) {
	data := gin.H{
		"Title": "Politique de confidentialité",
	}

	ShowPage(c, "policy", data)
}