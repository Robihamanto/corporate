package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

// JwtGroup is list of endpoint of Jwt
func JwtGroup(g *echo.Group) {
	g.GET("/main", handlers.GetJwtHome)
}
