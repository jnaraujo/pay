package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jnaraujo/pay/internal/services"
	"github.com/jnaraujo/pay/internal/validate"
)

type CreateGameSchema struct {
	OwnerId uuid.UUID `json:"owner_id" validate:"required"`
}

func CreateGameRoute(c *fiber.Ctx) error {
	var body CreateGameSchema
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if errs := validate.Validate(body); len(errs) > 0 && errs[0].Error {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errs})
	}

	game, err := services.CreateGame(body.OwnerId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(game)
}
