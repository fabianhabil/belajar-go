package server

import (
	"belajar-go/config"
	"belajar-go/routes"	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func Serve() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	routes.RouteV1(app)

	log.Fatal(app.Listen(":" + config.PORT))
}
