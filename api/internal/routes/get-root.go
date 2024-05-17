package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetRoot(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("OK")
}
