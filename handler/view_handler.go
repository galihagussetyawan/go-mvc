package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func HomeView(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Go Mvc",
	})
}

func RegisterView(c *fiber.Ctx) error {
	tokenString := c.Cookies("token")

	if tokenString != "" {
		token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret-token-key-hardcoded"), nil //harcoded secret-key
		})

		if token.Valid {
			return c.Redirect("/")
		}
	}

	return c.Render("register", fiber.Map{})
}

func ProfileView(c *fiber.Ctx) error {
	return c.Render("profile", fiber.Map{})
}

func AdminView(c *fiber.Ctx) error {
	return c.Render("admin", fiber.Map{})
}
