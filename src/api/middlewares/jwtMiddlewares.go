package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//SetJwtMiddlewares create middleware that will be used in router
func SetJwtMiddlewares(g *echo.Group) {
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	// g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningMethod: "HS512",
	// 	SigningKey:    []byte("asDf#$#!@#0~!o"),
	// 	TokenLookup:   "cookie:JWTCookie",
	// }))

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("asDf#$#!@#0~!o"),
		TokenLookup:   "header:Authorization",
		AuthScheme:    "Bearer",
	}))
}
