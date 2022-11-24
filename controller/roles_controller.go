package controller

import (
	"log"
	"workspace_booking/utility"

	"workspace_booking/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// AllRoles from db
func AllRoles(c *fiber.Ctx) error {
	// query role table in the model

	roles := model.GetAllRoles()
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if err := c.JSON(&fiber.Map{
		"success": true,
		"role":    roles,
		"user":    claims,
		"message": "All role returned successfully",
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting roles", 500, err)
	}
	return nil
}

// CreateRole handler
func CreateRole(c *fiber.Ctx) error {
	// Instantiate new Role struct
	r := new(model.Role)

	//  Parse body into role struct
	if err := c.BodyParser(r); err != nil {
		return utility.ErrResponse(c, "Error in parsing", 400, err)
	}
	err := r.InsertRole()

	if err != nil {
		return utility.ErrResponse(c, "Error in saving", 500, err)
	}

	// Print result
	log.Println(r)

	// Return Role in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"role":    r,
		"message": "Role successfully created",
	}); err != nil {
		return utility.ErrResponse(c, "Error in sending response", 500, err)
	}
	return nil
}
