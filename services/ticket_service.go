package services

import (
	"errors"

	"github.com/marcelfc/ticketing-system/models"
	"github.com/marcelfc/ticketing-system/repositories"
)

type TicketService interface {
	BookTicket(eventID uint, userID uint) (*models.Ticket, error)
	GetTickets() ([]models.Ticket, error)
	GetTicketsByEventID(eventID uint) ([]models.Ticket, error)
}

type ticketService struct {
	ticketRepo repositories.TicketRepository
	eventRepo	repositories.EventRepository
}

func NewTicketService(ticketRepo repositories.TicketRepository, eventRepo repositories.EventRepository) TicketService {
	return &ticketService{ticketRepo, eventRepo}
}

func (s *ticketService) BookTicket(eventID uint, userID uint) (*models.Ticket, error) {
	event, err := s.eventRepo.FindByID(eventID)
	if err != nil {
		return nil, err
	}
	if event.BookedSpots >= event.TotalSpots {
		return nil, errors.New("no spots available")
	}

	ticket := &models.Ticket{
		EventID: eventID,
		UserID:  userID,
	}

	err = s.ticketRepo.Create(ticket)

	if err != nil {
		return nil, err
	}

	event.BookedSpots++
	err = s.eventRepo.Update(event)

	if err != nil {
		return nil, err
	}

	return ticket, err
}

func (s *ticketService) GetTickets() ([]models.Ticket, error) {
	return s.ticketRepo.FindAll()
}

func (s *ticketService) GetTicketsByEventID(eventID uint) ([]models.Ticket, error) {
	return s.ticketRepo.FindByEventID(eventID)
}