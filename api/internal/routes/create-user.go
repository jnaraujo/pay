package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jnaraujo/pay/internal/services"
	"github.com/jnaraujo/pay/internal/validate"
)

type CreateUserSchema struct {
	Name string `json:"name" validate:"required"`
}

func CreateUserRoute(c *fiber.Ctx) error {
	var body CreateUserSchema
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if errs := validate.Validate(body); len(errs) > 0 && errs[0].Error {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errs})
	}

	createdUser, err := services.CreateUser(body.Name)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(createdUser)
}
