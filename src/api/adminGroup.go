package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

// AdminGroup is list of endpoint of Admin
func AdminGroup(g *echo.Group) {
	g.GET("/", handlers.HomeAdmin, handlers.CheckCookie)
	g.GET("/login", handlers.LoginAdmin)
}
