package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name string
	Description string
	TotalSpots int
	BookedSpots int
}