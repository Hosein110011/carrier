package router

import (
	"carrier/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Carrier
	carrier := api.Group("carrier")
	carrier.Post("/", controller.OrderHandler)
}
