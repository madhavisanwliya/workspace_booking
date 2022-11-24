package controller

import (
	"strconv"
	m "workspace_booking/model"
	"workspace_booking/utility"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	u := new(m.User)

	if err := c.BodyParser(u); err != nil {
		return utility.ErrResponse(c, "Error in parsing", 400, err)
	}

	err := u.InsertUser()
	if err != nil {
		return utility.ErrResponse(c, "Error in creating user", 400, err)
	}

	return c.JSON(fiber.Map{
		"message": "User Successfully Added",
		"user":    u,
	})
}

func GetUsers(c *fiber.Ctx) error {
	users := new(m.Users)

	users, err := users.FetchUsers()

	if err != nil {
		return utility.ErrResponse(c, "Error in fetching users", 400, err)
	}

	return c.JSON(fiber.Map{
		"users": users,
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	i, e := strconv.Atoi(id)

	if e != nil {
		return c.Status(400).SendString(e.Error())
	}

	u := m.User{ID: i}
	err := u.FetchUser()

	// user, err := m.FetchUser(db, id)
	if err != nil {
		return utility.ErrResponse(c, "Error in fetching user", 400, err)
	}

	return c.JSON(fiber.Map{
		"user": u,
	})
}

func EditUser(c *fiber.Ctx) error {
	id := c.Params("id")

	i, e := strconv.Atoi(id)

	if e != nil {
		return c.Status(400).SendString(e.Error())
	}
	u := &m.User{ID: i}

	if err := c.BodyParser(u); err != nil {
		return utility.ErrResponse(c, "Error in parsing", 400, err)
	}
	err := u.UpdateUser()

	if err != nil {
		return utility.ErrResponse(c, "Error in updating user", 400, err)
	}

	return c.JSON(fiber.Map{
		"message": "User Successfully Updated",
		"user":    u,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	i, _ := strconv.Atoi(id)

	u := &m.User{ID: i}

	err := u.DeleteUser()

	if err != nil {
		return utility.ErrResponse(c, "Error in deleting user", 400, err)
	}

	return c.JSON(fiber.Map{
		"message": "User Successfully Deleted",
	})
}
