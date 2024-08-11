package repositories

import (
	"github.com/marcelfc/ticketing-system/models"
	"gorm.io/gorm"
)

type EventRepository interface {
	Create(event *models.Event) error
	FindAll() ([]models.Event, error)
	FindByID(id uint) (*models.Event, error)
	Update(event *models.Event) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) Create(event *models.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) FindAll() ([]models.Event, error) {
	var events []models.Event
	err := r.db.Find(&events).Error
	return events, err
}

func (r *eventRepository) FindByID(id uint) (*models.Event, error) {
	var event models.Event
	err := r.db.First(&event, id).Error
	return &event, err
}

func (r *eventRepository) Update(event *models.Event) error {
	return r.db.Save(event).Error
}