package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// HTML Location
	r.LoadHTMLFiles("html/index.html")

	r.StaticFile("/logo-light.png", "./assets/logo-light.png")
	r.StaticFile("/style.css", "./style/style.css")
	r.StaticFile("/bulma-carousel.min.css", "./bulma-carousel/dist/css/bulma-carousel.min.css")
	r.StaticFile("/bulma-carousel.min.js", "./bulma-carousel/dist/js/bulma-carousel.min.js")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Run(":8081")
}
