package handler

import (
	"github.com/gofiber/fiber/v2"
)

func HomeView(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Go Mvc",
	})
}

func RegisterView(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{})
}
