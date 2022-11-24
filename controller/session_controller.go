package controller

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"workspace_booking/config"
	"workspace_booking/mailer"
	m "workspace_booking/model"
	"workspace_booking/utility"

	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return utility.ErrResponse(c, "Error in parsing", 400, err)
	}
	log.Println(data)

	log.Println([]byte(data["password"]))

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 10)

	u := &m.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	log.Println("pss", password)

	errors := utility.ValidateUserStruct(*u)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	emailDomain := strings.Split(u.Email, "@")[1]

	if emailDomain != config.GetEmailDomain() {
		return utility.ErrResponse(c, "Invalid Email", 500, nil)
	}

	err := u.InsertUser()

	if err != nil || u.ID == 0 {
		return utility.ErrResponse(c, "User is already exist", 500, err)
	}

	if u.ID != 0 {
		go mailer.UserMailer(u)
	}

	log.Println("u.pss", u.Password)
	log.Println(u)
	return c.JSON(&fiber.Map{
		"message": "Successfully register",
		"user":    u,
	})
}

func Login(c *fiber.Ctx) error {
	// user := c.FormValue("user")
	// pass := c.FormValue("pass")

	// Throws Unauthorized error
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return utility.ErrResponse(c, "Error in parsing", 400, err)
	}

	u := new(m.User)

	if err := c.BodyParser(u); err != nil {
		return utility.ErrResponse(c, "Error in parsing", 400, err)
	}

	err := u.LoginUser()

	if err != nil {
		return utility.ErrResponse(c, "Invalid Access!", 401, err)
	}

	if !(u.VerifiedUser) {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not verified please verify",
		})
	}
	if u.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(data["password"]))

	if err != nil {
		return utility.ErrResponse(c, "Incorrect Password", 400, err)
	}

	uID := strconv.Itoa(u.ID)

	// Create the Claims
	claims := jwt.MapClaims{
		"id":    uID,
		"name":  u.Name,
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.GetJWTSecret()))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:    u.Email,
		Value:   t,
		Expires: time.Now().Add(24 * time.Hour),
	})

	return c.JSON(fiber.Map{
		"message": "Successfully login",
		"token":   t,
		"user":    u,
	})
}

func Logout(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	log.Println(email)
	c.Cookie(&fiber.Cookie{
		Name:    email,
		Value:   "",
		Expires: time.Now().Add(5 * time.Second),
	})
	return c.JSON(fiber.Map{
		"message": "Successfully logout",
	})
}

func VerifyOtp(c *fiber.Ctx) error {
	u := new(m.User)

	if err := c.BodyParser(u); err != nil {
		return utility.ErrResponse(c, "Somthing went wrong", 400, err)
	}

	err := u.VerifyUser()

	if err != nil {
		return utility.ErrResponse(c, "Invalid OTP", 401, err)
	}

	if u.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Invalid OTP",
		})
	}
	err = u.UpdateVerifyUser()

	if err != nil {
		return utility.ErrResponse(c, "Somthing went wrong while updating", 401, err)
	}
	return c.JSON(fiber.Map{
		"message": "Successfully verified",
	})
}
