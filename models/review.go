package models

import (
	"time"
)

type Review struct {
	ReviewId  string `gorm:"unique"`
	MovieId   int
	Stars     int    `validate:"required"`
	Title     string `validate:"required"`
	Content   string `validate:"required"`
	CreatedAt time.Time
}
