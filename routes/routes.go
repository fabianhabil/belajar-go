package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouteV1(app *fiber.App) {
	v1 := app.Group("/v1")
	UserRoute(v1)
}
