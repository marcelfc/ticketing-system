package services

import (
	"github.com/marcelfc/ticketing-system/models"
	"github.com/marcelfc/ticketing-system/repositories"
)

type EventService interface {
	CreateEvent(name string, description string, totalSpots int) (*models.Event, error)
	GetEvents() ([]models.Event, error)
	GetEventByID(id uint) (*models.Event, error)
	UpdateEvent(id uint, name string, description string, totalSpots int) (*models.Event, error)
}

type eventService struct {
	eventRepo repositories.EventRepository
}

func NewEventService(eventRepo repositories.EventRepository) EventService {
	return &eventService{eventRepo}
}

func (s *eventService) CreateEvent(name string, description string, totalSpots int) (*models.Event, error) {
	event := &models.Event{
		Name:        name,
		Description: description,
		TotalSpots:  totalSpots,
	}
	err := s.eventRepo.Create(event)
	return event, err
}

func (s *eventService) GetEvents() ([]models.Event, error) {
	return s.eventRepo.FindAll()
}

func (s *eventService) GetEventByID(id uint) (*models.Event, error) {
	return s.eventRepo.FindByID(id)
}

func (s *eventService) UpdateEvent(id uint, name string, description string, totalSpots int) (*models.Event, error) {
	event, err := s.eventRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	event.Name = name
	event.Description = description
	event.TotalSpots = totalSpots
	err = s.eventRepo.Update(event)
	return event, err
}