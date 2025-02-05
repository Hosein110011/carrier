package controller

import (
	"carrier/model/carrier"
	"carrier/model/destination"
	"carrier/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func OrderHandler(c *fiber.Ctx) error {
	destination := new(destination.Destination)
	if err := c.BodyParser(&destination); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Incorrect input", "errors": err.Error()})
	}
	var minDistance float64
	var selectedCarrier carrier.Carrier
	carriers, _ := service.GetAvailableCarriers()
	for index, carrier := range carriers {
		selectedCarrier = carrier
		distance := service.GetDistance(carrier.X, carrier.Y, destination.Xd, destination.Yd)
		if index == 0 {
			minDistance = distance
		}
		minDistance = service.FindMinDistance(minDistance, distance)
		fmt.Println("min distance == ", minDistance, "distance == ", distance)
	}
	fmt.Println("minimum total distance == ", minDistance)
	counterX := service.GetAverageCounter(selectedCarrier.X, destination.Xd, minDistance)
	counterY := service.GetAverageCounter(selectedCarrier.Y, destination.Yd, minDistance)
	fmt.Println("CX : ", counterX, " CY : ", counterY)
	if err := service.CarrierTaskMaker(counterX, counterY, minDistance, selectedCarrier); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "something went wrong!", "errors": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "a carrier found to carry your load."})
}
