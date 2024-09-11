package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CustomLogger(c *fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	// After the request has been processed
	duration := time.Since(start)
	path := c.Path()
	status := c.Response().StatusCode()

	// Log the request and response
	log.Printf("Method: %s, Path: %s, Status: %d, Duration: %v", c.Method(), path, status, duration)

	return err
}
