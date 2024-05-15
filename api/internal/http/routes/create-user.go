package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jnaraujo/pay/internal/models"
	"github.com/jnaraujo/pay/internal/services"
	"github.com/jnaraujo/pay/internal/validate"
)

func CreateUserRoute(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if errs := validate.Validate(user); len(errs) > 0 && errs[0].Error {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errs})
	}

	createdUser, err := services.CreateUser(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(createdUser)
}
