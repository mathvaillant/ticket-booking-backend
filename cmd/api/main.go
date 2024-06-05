package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mathvaillant/ticket-booking-project-v0/config"
	"github.com/mathvaillant/ticket-booking-project-v0/db"
	"github.com/mathvaillant/ticket-booking-project-v0/handlers"
	"github.com/mathvaillant/ticket-booking-project-v0/middlewares"
	"github.com/mathvaillant/ticket-booking-project-v0/repositories"
	"github.com/mathvaillant/ticket-booking-project-v0/services"
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
	authRepository := repositories.NewAuthRepository(db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
