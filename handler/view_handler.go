package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func HomeView(c *fiber.Ctx) error {

	log.Printf("email cookies: %v", c.Cookies("email"))
	return c.Render("index", fiber.Map{
		"Title": "Go Mvc",
	})
}

func RegisterView(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{})
}
