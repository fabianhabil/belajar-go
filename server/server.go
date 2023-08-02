package server

import (
	"belajar-go/config"
	"belajar-go/database"
	"belajar-go/database/migration"
	"belajar-go/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Serve() {
	database.InitDB()
	migration.DBMigration()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	routes.RouteV1(app)

	log.Fatal(app.Listen(":" + config.PORT))
}
