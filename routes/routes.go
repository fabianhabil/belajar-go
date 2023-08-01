package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouteV1(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})
}
