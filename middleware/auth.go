package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func jwtParser(t string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret-token-key-hardcoded"), nil //harcoded secret-key
	})

	return token, err
}

func LoginRequired(c *fiber.Ctx) error {
	tokenString := c.Cookies("token")

	if tokenString == "" {
		return c.Redirect("/")
	}

	token, err := jwtParser(tokenString)
	if err != nil {
		return c.Redirect("/")
	}

	if !token.Valid {
		return c.Redirect("/")
	}

	return c.Next()
}

func AdminRequired(c *fiber.Ctx) error {
	tokenString := c.Cookies("token")

	if tokenString == "" {
		return c.Redirect("/")
	}

	token, err := jwtParser(tokenString)
	if err != nil {
		return c.Redirect("/")
	}

	if !token.Valid {
		return c.Redirect("/")
	}

	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"]

	if role != "admin" {
		return c.Redirect("/")
	}

	return c.Next()
}
