package main

import (
	"github.com/gofiber/fiber/v2"
	"golangapi/api"
	"golang/kafka"
	"golang/utils"
)

func main() {

	app := fiber.New()
	utils.SetupLogger()

	kafka.InitializeProducer()
	kafka.InitializeConsumer()

	api.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
