package middleware

import (
	"fmt"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/models/response"
	respCode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	"github.com/gofiber/fiber/v2"
)

func CheckRank(minRank int) fiber.Handler {
	logger.WithField("method", "CheckRank")
	return func(c *fiber.Ctx) error {
		rank := c.Locals("rank").(int)
		if rank <= 0 {

			logger.Info("Invalid rank attempted access")
			return c.Status(401).JSON(response.SRE{
				Status:       "failed",
				ResponseCode: respCode.Unauthorized,
				Error:        "Invalid rank",
			})
		} else if rank < minRank {
			logger.Info(fmt.Sprintf("Insufficient rank attempted access. Required rank: %d, Admin rank: %d", minRank, rank))

			return c.Status(401).JSON(response.SRE{
				Status:       "failed",
				ResponseCode: respCode.Unauthorized,
				Error:        "Insufficient rank",
			})
		}
		return c.Next()
	}
}
