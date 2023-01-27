package router

import (
	"go-mvc/handler"
	"go-mvc/middleware"

	"github.com/gofiber/fiber/v2"
)

func ViewRouter(a *fiber.App) {
	a.Get("/", handler.HomeView)
	a.Get("/register", handler.RegisterView)
	a.Get("/profile", middleware.LoginRequired, handler.ProfileView)
	a.Get("/admin", middleware.AdminRequired, handler.AdminView)
}
