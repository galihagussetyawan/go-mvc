package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func LoginRequired(c *fiber.Ctx) error {
	tokenString := c.Cookies("token")

	if tokenString == "" {
		return c.Redirect("/")
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret-token-key-hardcoded"), nil //harcoded secret-key
	})

	if err != nil {
		return c.Redirect("/")
	}

	if !token.Valid {
		return c.Redirect("/")
	}

	return c.Next()
}
