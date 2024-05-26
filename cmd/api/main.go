package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mathvaillant/ticket-booking-project-v0/config"
	"github.com/mathvaillant/ticket-booking-project-v0/db"
	"github.com/mathvaillant/ticket-booking-project-v0/handlers"
	"github.com/mathvaillant/ticket-booking-project-v0/repositories"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Ticket-Booking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)

	// Routing
	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)
	handlers.NewTicketHandler(server.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
