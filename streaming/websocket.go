package streaming

import (
	"time"
	
	"github.com/gofiber/fiber/v2"
)

func StreamResults(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	for {
		c.Write([]byte("data: Your message here \n\n"))
		// c.Context().Flush()
		time.Sleep(1 * time.Second)
	}
}
