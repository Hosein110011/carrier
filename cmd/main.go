package main

import (
	"carrier/database"
	"carrier/router"
	"carrier/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Carriers",
	})

	database.DB = database.ConnectDB()

	router.SetupRoutes(app)

	err := service.InitializeCarriers()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Fatal(app.Listen(":3000"))
}
