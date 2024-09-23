package main

import "github.com/gin-gonic/gin"

func PanelAdminHandler(c *gin.Context) {
	// Get list of VM from Webmin API

	data := gin.H{
		"Title": "Administration",
	}

	ShowPage(c, "panel-admin", data)
}