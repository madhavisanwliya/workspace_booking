package controller

import (
	"workspace_booking/model"
	"workspace_booking/utility"

	"github.com/gofiber/fiber/v2"
)

// Index
func AllLocations(c *fiber.Ctx) error {
	locations := model.GetAllLocations()

	if err := c.JSON(&fiber.Map{
		"success":   true,
		"locations": locations,
		"message":   "All locations returned successfully",
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting locations", 500, err)
	}
	return nil
}

// CreateLocation handler
func CreateLocation(c *fiber.Ctx) error {

	location := new(model.Location)
	if err := c.BodyParser(location); err != nil {
		return utility.ErrResponse(c, "Error in creating location", 400, err)
	}

	err := location.CreateLocation()
	if err != nil {
		return utility.ErrResponse(c, "Error in creating location", 500, err)
	}

	// Return Location in JSON format
	if err := c.JSON(&fiber.Map{
		"success":  true,
		"location": location,
		"message":  "Location successfully created",
	}); err != nil {
		return utility.ErrResponse(c, "Error in creating location", 500, err)
	}
	return nil
}
