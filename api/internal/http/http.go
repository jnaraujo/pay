package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jnaraujo/pay/internal/config"
	"github.com/jnaraujo/pay/internal/http/routes"
)

func NewServer() error {
	app := fiber.New()

	registerRoutes(app)

	err := app.Listen(config.Env.ServerUrl)
	if err != nil {
		return err
	}
	return nil
}

func registerRoutes(app *fiber.App) {
	router := app.Group("/api/v1")
	router.Get("/", routes.GetRoot)
}
