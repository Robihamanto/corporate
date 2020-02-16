package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// User for define
type User struct {
	Name string "json:name"
	Age  int    "json:age"
}

// Address of users
type Address struct {
	User       int    "json:user"
	Street     string "json:street"
	PostalCode int    "json:postalCode"
}

// Car has owned by users
type Car struct {
	User  int    "json:user"
	Color string "json:color"
}

// ServerHeader Set header for middleware
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Staging")
		c.Response().Header().Set("Version", "1")
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Use(ServerHeader)
	adminGroup := e.Group("/admin")

	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	// adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	if username == "admin" && password == "admin" {
	// 		log.Println("Users successfully logged in")
	// 		return true, nil
	// 	}
	// 	log.Println("Username or password error")
	// 	return false, nil
	// }))

	//adminGroup.Use(checkCookie)
	adminGroup.GET("/", homeAdmin, checkCookie)
	adminGroup.GET("/login", loginAdmin)

	e.GET("/", landing)
	e.GET("/person", getUser)
	e.POST("/person", addUser)
	e.POST("/address", addAddress)
	e.POST("/car", addCar)

	e.Start(":8000")

}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
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

func loginAdmin(c echo.Context) error {
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

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You're logged in",
		})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{
		"error": "Username or password were wrong",
	})

}

func adminCookie(c echo.Context) error {
	return c.String(http.StatusOK, "you're not logged in")
}

func homeAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "You're logged in to scret cookie page",
	})
}

func getAddress(c echo.Context) error {
	userID := c.QueryParam("user")

	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Please provide user id",
		})
	}

	//select from db
	userFound := userID == "21"

	if userFound {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "There is no user with that id",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"user":    "21",
		"address": "Xianyi Street Number 38",
	})
}

func landing(c echo.Context) error {
	return c.String(http.StatusOK, "Hi this is the server run")
}

func getUser(c echo.Context) error {
	name := c.QueryParam("name")
	age := c.QueryParam("age")

	dataType := c.QueryParam("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintln("He/She is ", name, " and ", age, " years old"))
	} else if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": name,
			"age":  age,
		})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Undefined data type",
		})
	}
}

func addUser(c echo.Context) error {
	user := User{}
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Cannot find json body")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Internal server error",
		})
	}

	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Println("Cannot unmasrhal json body")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Internal server error, cannot unmarshal JSON",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User has been added",
	})
}

func addCar(c echo.Context) error {
	car := Car{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&car)
	if err != nil {
		log.Println("Car Address : Cannot find json body")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Internal server error, cannot unmarshal JSON",
		})
	}

	log.Println("New Car: ", car)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Car succesffully added",
	})
}

func addAddress(c echo.Context) error {
	address := Address{}

	err := c.Bind(&address)
	if err != nil {
		log.Println("Add Address : Cannot find json body")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Internal server error, cannot unmarshal JSON",
		})
	}

	log.Println("New Address: ", address)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User succesffully added",
	})

}
