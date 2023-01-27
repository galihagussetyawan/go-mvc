package router

import (
	"go-mvc/handler"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(a *fiber.App) {
	a.Post("/register", handler.Login)
}
