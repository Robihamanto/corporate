package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"model"
	"net/http"

	"github.com/labstack/echo"
)

// GetUser retrieving user data
func GetUser(c echo.Context) error {
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

// AddUser add user data
func AddUser(c echo.Context) error {
	user := model.User{}
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
