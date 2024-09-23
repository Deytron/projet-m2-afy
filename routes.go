package main

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine) {
	// Routes declaration

	// 404
	r.NoRoute(NoRouteHandler)

	// GET Routes
	r.GET("/", HomeHandler)
	r.GET("/login", LoginHandler)
	r.GET("/register", RegisterHandler)
	r.GET("/logout", LogoutHandler)
	r.GET("/unauthorized", UnauthorizedHandler)
	r.GET("/panel-admin", PanelAdminHandler)
	r.GET("/vms", VmsHandler)
	r.GET("/create-vm", CreateVmsHandler)
	r.GET("/policy", PolicyHandler)
	r.GET("/sites", SitesHandler)
	r.GET("/names", NamesHandler)
	r.GET("/support", SupportHandler)

	// POST routes
	r.POST("/login", LoginHandler)
	r.POST("/register", RegisterHandler)
	r.POST("/create-vm", CreateVmsHandler)
}
