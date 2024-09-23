package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var pageAuthorization = map[string]int{
	"categories":       4,
	"dc-logs":          4,
	"get-dhcp":         4,
	"add-dhcp":         2,
	"favorites":        1,
	"ftp-logs":         3,
	"home":             1,
	"huawei-mac":       4,
	"presta":           4,
	"pwsh-scriptslist": 2,
	"pwsh-script":      2,
	"users":            4,
	"edit-shortcut":    2,
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NoRouteHandler(c *gin.Context) {
	data := gin.H{
		"Title": "Page non trouvée",
	}

	ShowPage(c, "404", data)
}

func UnauthorizedHandler(c *gin.Context) {

	fmt.Println(pageAuthorization["categories"])
	data := gin.H{
		"Title": "Non autorisé",
	}

	ShowPage(c, "unauthorized", data)
}

func ShowPage(c *gin.Context, page string, data gin.H) {
	// Using templates
	var t = GetConfig().Templates

	if SuccessMessage != "" || ErrorMessage != "nil" {
		data["SuccessMessage"] = SuccessMessage
		data["ErrorMessage"] = ErrorMessage
	}

	_, cerr := c.Cookie("username") 
	if cerr == nil {
		data["Logged"] = true
	}

	err := t.ExecuteTemplate(c.Writer, page+".html", data)
	if err != nil {
		fmt.Println(500, err.Error())
	}

	SuccessMessage = ""
	ErrorMessage = ""
}
