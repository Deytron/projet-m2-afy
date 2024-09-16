package main

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

// Set variables for use across all files
type AppConfig struct {
	Templates *template.Template
	Router    *gin.Engine
}

var appConfig *AppConfig

func InitConfig() {
	// gin.SetMode(gin.ReleaseMode)
	appConfig = &AppConfig{
		Templates: template.Must(template.ParseGlob("/app/html/*.html")),
		Router:    gin.Default(),
	}
}

func GetConfig() *AppConfig {
	return appConfig
}
