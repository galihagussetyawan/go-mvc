package router

import (
	"go-mvc/handler"

	"github.com/gofiber/fiber/v2"
)

func ViewRouter(a *fiber.App) {
	a.Get("/", handler.HomeView)
	a.Get("/register", handler.RegisterView)
}
