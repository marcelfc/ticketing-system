package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/marcelfc/ticketing-system/services"
)

type TicketController struct {
	ticketService services.TicketService
}

func NewTicketController(ticketService services.TicketService) *TicketController {
	return &TicketController{ticketService}
}

func (tc *TicketController) BookTicket(c echo.Context) error {
	var input struct {
		EventID uint `json:"event_id" binding:"required"`
		UserID  uint `json:"user_id" binding:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ticket, err := tc.ticketService.BookTicket(input.EventID, input.UserID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, ticket)
}

func (tc *TicketController) GetTickets(c echo.Context) error {
	tickets, err := tc.ticketService.GetTickets()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, tickets)
}

func (tc *TicketController) GetTicketsByEventID(c echo.Context) error {
	eventID, err := strconv.Atoi(c.Param("event_id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid event id"})
	}

	tickets, err := tc.ticketService.GetTicketsByEventID(uint(eventID))

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "event not found"})
	}

	return c.JSON(http.StatusOK, tickets)
}