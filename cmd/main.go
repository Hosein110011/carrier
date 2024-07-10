package main

import (
	"carrier/database"
	"carrier/router"
	"carrier/service"
	"fmt"
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

	database.ConnectDB()

	if err := service.InitializeCarriers(); err != nil {
		fmt.Println("erroooooooor")
		panic(err.Error())
		return
	}

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
