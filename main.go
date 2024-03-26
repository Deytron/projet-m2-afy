package main

import (
	"net/http"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

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
