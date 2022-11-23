package router

import (
	"strings"
	"workspace_booking/config"
	"workspace_booking/controller"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	// No Auth routes
	api.Post("/sign-up", controller.Register)
	api.Post("/sign-in", controller.Login)
	api.Post("/verify-otp", controller.VerifyOtp)

	api.Get("/roles", controller.AllRoles)
	api.Post("/roles", controller.CreateRole)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.GetJWTSecret()),
	}))

	app.Use(func(c *fiber.Ctx) error {
		url_token := c.Get("Authorization")
		u_token := strings.Split(url_token, " ")[1]
		token, err := jwt.Parse(u_token, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetJWTSecret()), nil
		})
		if err != nil {
			return c.JSON(fiber.Map{
				"message": "Invalid Access",
			})
		}
		if token.Raw == u_token {
			c.Locals("verify", "true")
		} else {
			c.Locals("verify", "false")
		}
		return c.Next()
	})

	// Authorization routes

	api.Post("/users", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.CreateUser(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/users", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.GetUsers(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/users/:id", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.GetUser(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Delete("/users/:id", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.DeleteUser(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/logout", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.Logout(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Post("/book_workspace", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.CreateBooking(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/workspace_details", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.WorkSpacesDetails(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	/* Building API's */
	api.Post("/buildings", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.CreateBuilding(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/buildings", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.AllBuildings(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	/* City API's */
	api.Post("/cities", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.CreateCity(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/cities", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.AllCities(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	/* Location API's */
	api.Post("/locations", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.CreateLocation(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/locations", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.AllLocations(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	/* Floor API's */
	api.Post("/floors", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.CreateFloor(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/floors", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.AllFloors(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	/* available workspaces count*/
	api.Get("/available_workspace", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.GetAvailableBookingSpace(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/my_bookings", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.MyBookingDetails(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Get("/all_amenities", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.AllAmenities(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	api.Put("/cancel_booking/:id", func(c *fiber.Ctx) error {
		user := c.Locals("verify")
		if user == "true" {
			return controller.CanceledBookings(c)
		}
		return c.SendStatus(fiber.StatusForbidden)
	})
}
