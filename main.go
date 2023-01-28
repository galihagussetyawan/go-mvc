package main

import (
	"go-mvc/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/handlebars"
)

func main() {
	engine := handlebars.New("./views", ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: "wN9Uxz+6RFvUI/J/ir1zYIa1MhrNVi1NbV8J+bd5ojw=", //hardcoded key
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// app middlewares
	app.Use(logger.New())

	//router
	router.ViewRouter(app)
	router.AuthRouter(app)

	log.Fatal(app.Listen(":3000"))
}
