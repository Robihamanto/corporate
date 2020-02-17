package handlers

import (
	"encoding/json"
	"log"
	"model"
	"net/http"

	"github.com/labstack/echo"
)

// AddCar is adding car data to user data
func AddCar(c echo.Context) error {
	car := model.Car{}
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
