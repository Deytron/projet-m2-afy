package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SetSessionCookie sets a session cookie for the user
func SetSessionCookie(c *gin.Context, username string) {
	token := uuid.NewString()
	// WARNING If changing URL for site, also change HOST in .env
	c.SetCookie("session_token", token, 14400, "/", os.Getenv("HOST"), true, true)
	c.SetCookie("username", username, 14400, "/", os.Getenv("HOST"), true, true)
}

func LogoutHandler(c *gin.Context) {
	c.SetCookie("session_token", "XXX", -1000, "/", os.Getenv("HOST"), true, true)
	c.SetCookie("username", "XXX", -1000, "/", os.Getenv("HOST"), true, true)
	c.Redirect(http.StatusTemporaryRedirect, "/login")
}
