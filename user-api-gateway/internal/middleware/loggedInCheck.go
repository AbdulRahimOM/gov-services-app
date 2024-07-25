package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func NotLoggedInCheck(c *fiber.Ctx) error {
	return c.Next()
}
