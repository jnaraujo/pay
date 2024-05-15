package routes

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/jnaraujo/pay/internal/errs"
	"github.com/jnaraujo/pay/internal/services"
	"github.com/jnaraujo/pay/internal/validate"
)

type transferPaymentSchema struct {
	SenderId   uuid.UUID `json:"sender_id" validate:"required"`
	ReceiverId uuid.UUID `json:"receiver_id" validate:"required"`
	Amount     int       `json:"amount" validate:"required"`
}

func TransferPaymentRoute(c *fiber.Ctx) error {
	var body transferPaymentSchema
	if errs := validate.ParseAndValidate(c.Body(), &body); len(errs) > 0 && errs[0].Error {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errs})
	}

	if body.Amount <= 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Amount must be greater than 0",
		})
	}

	sender, err := services.FindUserById(body.SenderId)
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Sender not found",
		})
	}

	receiver, err := services.FindUserById(body.ReceiverId)
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Receiver not found",
		})
	}

	if sender.GameId == nil || receiver.GameId == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Sender and receiver must be in a game",
		})
	}

	if sender.Id == receiver.Id {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Sender and receiver must be different",
		})
	}

	if sender.GameId != receiver.GameId {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Sender and receiver must be in the same game",
		})
	}

	err = services.TransferPayment(body.SenderId, body.ReceiverId, body.Amount)
	if err != nil {
		if errors.Is(err, errs.ErrInsufficientBalance) {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Insufficient balance",
			})
		}

		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Payment transferred successfully",
	})
}
