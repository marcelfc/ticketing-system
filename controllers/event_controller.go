package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/marcelfc/ticketing-system/services"
)

type EventController struct {
	eventService services.EventService
}

func NewEventController(eventService services.EventService) *EventController {
	return &EventController{eventService}
}

func (ec *EventController) CreateEvent(c echo.Context) error {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		TotalSpots  int    `json:"total_spots" binding:"required"`
	}

	// check if the request body is valid
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	event, err := ec.eventService.CreateEvent(input.Name, input.Description, input.TotalSpots)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, event)
}

func (ec *EventController) GetEvents(c echo.Context) error {
	events, err := ec.eventService.GetEvents()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, events)
}

func (ec *EventController) GetEventByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid event id"})
	}

	event, err := ec.eventService.GetEventByID(uint(id))

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "event not found"})
	}

	return c.JSON(http.StatusOK, event)

}

func (ec *EventController) UpdateEvent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid event id"})
	}

	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		TotalSpots  int    `json:"total_spots"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	event, err := ec.eventService.UpdateEvent(uint(id), input.Name, input.Description, input.TotalSpots)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, event)
}