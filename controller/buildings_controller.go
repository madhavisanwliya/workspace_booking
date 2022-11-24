package controller

import (
	"workspace_booking/model"
	"workspace_booking/utility"

	"github.com/gofiber/fiber/v2"
)

// AllBuildings
func AllBuildings(c *fiber.Ctx) error {

	buildings := model.GetAllBuildings()
	if err := c.JSON(&fiber.Map{
		"success":   true,
		"buildings": buildings,
		"message":   "All Buildings returned successfully",
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting buildings", 500, err)
	}
	return nil
}

// CreateBuilding handler
func CreateBuilding(c *fiber.Ctx) error {

	building := new(model.Building)
	if err := c.BodyParser(building); err != nil {
		return utility.ErrResponse(c, "Error in creating building", 400, err)
	}

	err := building.CreateBuilding()
	if err != nil {
		return utility.ErrResponse(c, "Error in creating building", 500, err)
	}

	// Return Building in JSON format
	if err := c.JSON(&fiber.Map{
		"success":  true,
		"building": building,
		"message":  "Building successfully created",
	}); err != nil {
		return utility.ErrResponse(c, "Error in creating building", 500, err)
	}
	return nil
}
