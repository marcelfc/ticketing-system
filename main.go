package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/marcelfc/ticketing-system/controllers"
	"github.com/marcelfc/ticketing-system/database"
	"github.com/marcelfc/ticketing-system/models"
	"github.com/marcelfc/ticketing-system/repositories"
	"github.com/marcelfc/ticketing-system/routes"
	"github.com/marcelfc/ticketing-system/services"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
			log.Fatalf("Error loading .env file")
	}
			
	database.Connect()
	database.DB.AutoMigrate(&models.Event{}, &models.Ticket{})

	eventRepo := repositories.NewEventRepository(database.DB)
	ticketRepo := repositories.NewTicketRepository(database.DB)

	eventService := services.NewEventService(eventRepo)
	ticketService := services.NewTicketService(ticketRepo, eventRepo)

	eventController := controllers.NewEventController(eventService)
	ticketController := controllers.NewTicketController(ticketService)

	e := routes.SetupRouter(eventController, ticketController)

	e.Logger.Fatal(e.Start(":8080"))
}