package controller

import (
	"workspace_booking/model"
	"workspace_booking/utility"

	"github.com/gofiber/fiber/v2"
)

// AllCities
func AllCities(c *fiber.Ctx) error {

	cities := model.GetAllCities()
	if err := c.JSON(&fiber.Map{
		"success": true,
		"cities":  cities,
		"message": "All Cities returned successfully",
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting cities", 500, err)
	}
	return nil
}

// CreateCity handler
func CreateCity(c *fiber.Ctx) error {

	city := new(model.City)
	if err := c.BodyParser(city); err != nil {
		return utility.ErrResponse(c, "Error in creating city", 400, err)
	}

	err := city.CreateCity()
	if err != nil {
		return utility.ErrResponse(c, "Error in creating city", 500, err)
	}

	// Return city in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"city":    city,
		"message": "City successfully created",
	}); err != nil {
		return utility.ErrResponse(c, "Error in creating city", 500, err)
	}
	return nil
}
