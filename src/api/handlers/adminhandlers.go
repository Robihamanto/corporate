package handlers

import (
	"log"
	"model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//LoginAdmin is a way to logged in to system
func LoginAdmin(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// check account from db
	if username == "admin" && password == "admin" {
		// set login cookie
		cookie := &http.Cookie{}
		cookie.Name = "SessionID"
		cookie.Value = "token"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)

		token, err := createJwtToken()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "something happened in our server :(",
			})
		}

		// set JWT login cookie
		cookieJwt := &http.Cookie{}
		cookieJwt.Name = "JWTCookie"
		cookieJwt.Value = token
		cookieJwt.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookieJwt)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You're logged in",
			"token":   token,
		})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{
		"error": "Username or password were wrong",
	})

}

func adminCookie(c echo.Context) error {
	return c.String(http.StatusOK, "you're not logged in")
}

//HomeAdmin is retrieving home page
func HomeAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "You're logged in to scret cookie page",
	})
}

func createJwtToken() (string, error) {
	claims := model.JwtClaims{
		"anwar",
		jwt.StandardClaims{
			Id:        "user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte("asDf#$#!@#0~!o"))
	if err != nil {
		log.Print("failed generating JWT token")
		return "", err
	}
	return token, nil
}

// CheckCookie is checking is cookie saved in local
func CheckCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("SessionID")
		if err != nil {
			log.Println("Check cookies: ", err)
			return c.String(http.StatusUnauthorized, "You have to login first")
		}

		if cookie.Value == "token" {
			log.Println("Token is valid")
			return next(c)
		}

		log.Println("Token expired")
		return c.String(http.StatusUnauthorized, "You don't have any cookie or cookie has beed expired")
	}
}
