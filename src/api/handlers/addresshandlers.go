package handlers

import (
	"log"
	"model"
	"net/http"

	"github.com/labstack/echo"
)

// GetAddress retrieving user address
func GetAddress(c echo.Context) error {
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

// AddAddress retrieving user address
func AddAddress(c echo.Context) error {
	address := model.Address{}

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
