package controller

import (
	"carrier/model/carrier"
	"carrier/model/destination"
	"carrier/service"
	"github.com/gofiber/fiber/v2"
)

func OrderHandler(c *fiber.Ctx) error {
	destination := new(destination.Destination)
	if err := c.BodyParser(&destination); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Incorrect input", "errors": err.Error()})
	}
	var minDistance float64 = 0.0
	var selectedCarrier carrier.Carrier
	carriers, _ := service.GetAvailableCarriers()
	for _, carrier := range carriers {
		selectedCarrier = carrier
		distance := service.DistanceMeter(carrier.X, carrier.Y, destination.Xd, destination.Yd)
		minDistance = service.MinFinder(minDistance, distance)
	}
	counterX := service.AverageCounterFinder(selectedCarrier.X, destination.Xd, minDistance)
	counterY := service.AverageCounterFinder(selectedCarrier.Y, destination.Yd, minDistance)
	service.CarrierTaskMaker(counterX, counterY)
}
