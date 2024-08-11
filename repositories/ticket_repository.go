package repositories

import (
	"github.com/marcelfc/ticketing-system/models"
	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ticket *models.Ticket) error
	FindAll() ([]models.Ticket, error)
	FindByEventID(eventID uint) ([]models.Ticket, error)
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db}
}

func (r *ticketRepository) Create(ticket *models.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) FindAll() ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := r.db.Find(&tickets).Error
	return tickets, err
}

func (r *ticketRepository) FindByEventID(eventID uint) ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := r.db.Where("event_id = ?", eventID).Find(&tickets).Error
	return tickets, err
}