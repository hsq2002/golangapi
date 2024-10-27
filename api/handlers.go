package api

import (
	"github.com/gofiber/fiber/v2"

	"golangapi/streaming"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1", apiKeyMiddleware)
	api.Post("/stream/start", startStream)
	api.Post("/stream/:stream_id/send", sendData)
	api.Get("/stream/:stream_id/results", getResults)
	api.Get("stream/:stream_id/stream", streaming.StreamResults)
}

func startStream(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": true})
}

func sendData(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": true})
}


func getResults(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": true})
}
