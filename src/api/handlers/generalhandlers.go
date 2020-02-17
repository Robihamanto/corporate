package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

//HomePage is retrieving home page
func HomePage(c echo.Context) error {
	return c.String(http.StatusOK, "Hi this is the server run")
}

//Landing is retrieving home page
func Landing(c echo.Context) error {
	return c.String(http.StatusOK, "Hi this is the server run")
}
