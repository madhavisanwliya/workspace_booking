package controller

import (
	"workspace_booking/model"
	"workspace_booking/utility"

	"github.com/gofiber/fiber/v2"
)

// AllFloors
func AllFloors(c *fiber.Ctx) error {
	floors := model.GetAllFloors()

	if err := c.JSON(&fiber.Map{
		"success": true,
		"floors":  floors,
		"message": "All Floors returned successfully",
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting floors", 500, err)
	}
	return nil
}

// CreateFloor handler
func CreateFloor(c *fiber.Ctx) error {

	floor := new(model.Floor)
	if err := c.BodyParser(floor); err != nil {
		return utility.ErrResponse(c, "Error in creating floor", 400, err)
	}
	err := floor.CreateFloor()
	if err != nil {
		return utility.ErrResponse(c, "Error in creating floor", 500, err)
	}

	floorWorkspaceRecords := model.BulkFloorWorkspacesCreate(floor.Id, floor.FloorWorkSpaces)
	floor.FloorWorkSpaces = floorWorkspaceRecords

	// Return Floor in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"floor":   floor,
		"message": "Floor successfully created",
	}); err != nil {
		return utility.ErrResponse(c, "Error in creating floor", 500, err)
	}
	return nil
}
