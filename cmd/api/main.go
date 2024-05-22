package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mathvaillant/ticket-booking-project-v0/handlers"
	"github.com/mathvaillant/ticket-booking-project-v0/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(nil)

	// Routing
	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}
