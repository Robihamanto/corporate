package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// SetMainMiddlewares creating middleware for main router
func SetMainMiddlewares(e *echo.Echo) {
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "../static",
	}))

	e.Use(serverHeader)
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Staging")
		c.Response().Header().Set("Version", "1.0")
		return next(c)
	}
}
