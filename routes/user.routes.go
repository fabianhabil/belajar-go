package routes

import (
	"belajar-go/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(v1 fiber.Router) {
	user := v1.Group("/user")

	user.Get("/", controller.GetAllUser)
	user.Post("/create", controller.CreateUser)
}
