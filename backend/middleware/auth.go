package middleware

import (
	"github.com/JingusJohn/go-angular-twiddit/backend/api"
	"github.com/gofiber/fiber/v2"
)

var ()

func NewMiddleware(a func(*fiber.Ctx) error) fiber.Handler {
	return a
}

func AuthRequired(c *fiber.Ctx) error {
	session, err := api.SessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	if session.Get("authenticated") == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	return c.Next()

}
