package config

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AuthDetails struct {
	UserID    string
	UserEmail string
	UserName  string
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GetAuthDetails(c *fiber.Ctx) (*AuthDetails, error) {
	url_token := c.Get("Authorization")
	u_token := strings.Split(url_token, " ")[1]
	token, err := jwt.Parse(u_token, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetJWTSecret()), nil
	})

	if err != nil {
		return nil, c.JSON(fiber.Map{
			"message": "Invalid Access",
		})
	}
	claims, _ := token.Claims.(jwt.MapClaims)

	return &AuthDetails{
		UserID:    claims["id"].(string),
		UserEmail: claims["email"].(string),
		UserName:  claims["name"].(string),
	}, nil
}
