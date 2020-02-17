package handlers

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// GetJwtHome is retrieving Home page data with JWT Token
func GetJwtHome(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	log.Println("Username: ", claims["name"], "User ID: ", claims["jti"])

	return c.String(http.StatusOK, "You're in jwt secret home")
}
