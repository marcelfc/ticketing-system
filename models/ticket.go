package models

import "gorm.io/gorm"

type TicketTypeOption string

const (
	TicketTypeFree TicketTypeOption = "free"
	TicketTypePaid  TicketTypeOption = "paid"
)

type Ticket struct {
	gorm.Model
	EventID uint
	UserID uint
	TicketType TicketTypeOption
}