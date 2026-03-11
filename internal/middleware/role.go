package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RequireRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userRole := c.Locals("role")

		if userRole == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}

		if !strings.EqualFold(fmt.Sprintf("%v", userRole), role) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "forbidden",
			})
		}

		return c.Next()
	}
}
