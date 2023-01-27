package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	users := []string{"galih", "agus", "setyawan"}

	if email == "" || password == "" {
		return c.Render("register", fiber.Map{
			"STATUS": "email dan password harus diisi",
			"users":  users,
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "email",
		Value:    email,
		Secure:   true,
		HTTPOnly: true,
	})

	return c.Redirect("/")
}
