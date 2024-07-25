package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ClearCache(c *fiber.Ctx) error {
	fmt.Println("------ClearCache------")
	c.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	c.Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	return c.Next()
}
