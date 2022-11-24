package controller

import (
	"workspace_booking/model"
	"workspace_booking/utility"

	"github.com/gofiber/fiber/v2"
)

// AllAmenities
func AllAmenities(c *fiber.Ctx) error {
	amenities := model.GetAllAmenities()

	if err := c.JSON(&fiber.Map{
		"success":   true,
		"amenities": amenities,
		"message":   "All Amenities returned successfully",
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting Amenities", 500, err)
	}
	return nil
}

// CreateAmenity handler
func CreateAmenity(c *fiber.Ctx) error {

	amenity := new(model.Amenity)
	if err := c.BodyParser(amenity); err != nil {
		return utility.ErrResponse(c, "Error in creating amenity", 400, err)
	}
	err := amenity.CreateAmenity()
	if err != nil {
		return utility.ErrResponse(c, "Error in creating amenity", 500, err)
	}

	// Return Amenity in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"amenity": amenity,
		"message": "Amenity successfully created",
	}); err != nil {
		return utility.ErrResponse(c, "Error in creating amenity", 500, err)
	}
	return nil
}
