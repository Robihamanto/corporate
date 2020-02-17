package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

// MainGroup is list of endpoint of main
func MainGroup(e *echo.Echo) {
	e.GET("/person", handlers.GetUser)
	e.POST("/person", handlers.AddUser)
	e.POST("/address", handlers.AddAddress)
	e.POST("/car", handlers.AddCar)
	e.POST("/home", handlers.HomePage)
	e.POST("/landing", handlers.Landing)
}
