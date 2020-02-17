package router

import (
	"api"
	"api/middlewares"

	"github.com/labstack/echo"
)

// New function fot creating new router
func New() *echo.Echo {
	e := echo.New()

	adminGroup := e.Group("/admin")
	jwtGroup := e.Group("/jwt")

	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)

	api.MainGroup(e)
	api.JwtGroup(jwtGroup)
	api.AdminGroup(adminGroup)
	return e
}
