package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/marcelfc/ticketing-system/controllers"
)

func SetupRouter(eventController *controllers.EventController, ticketController *controllers.TicketController) *echo.Echo {

	e := echo.New()
	// Event routes
	e.GET("/events", eventController.GetEvents)
	e.POST("/events", eventController.CreateEvent)
	e.GET("/events/:id", eventController.GetEventByID)
	e.PUT("/events/:id", eventController.UpdateEvent)

	// Ticket routes
	e.POST("/tickets", ticketController.BookTicket)
	e.GET("/tickets", ticketController.GetTickets)
	e.GET("/tickets/event/:event_id", ticketController.GetTicketsByEventID)

	return e
}