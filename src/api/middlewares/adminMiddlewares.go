package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//SetAdminMiddlewares create middleware that will be used in router
func SetAdminMiddlewares(g *echo.Group) {
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
}
