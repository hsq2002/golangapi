package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"golangapi/api"
	"golangapi/kafka"
	"golangapi/utils"
)

func main() {
	app := fiber.New()
	utils.SetupLogger()

	brokersUrl := []string{"localhost:29092"}
	if err := kafka.InitializeProducer(brokersUrl); err != nil {
		log.Fatalf("Failed to initialize producer: %v", err)
	}

	if err := kafka.InitializeConsumer(brokersUrl); err != nil {
		log.Fatalf("Failed to initialize consumer: %v", err)
	}

	api.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
