package models

import (
	"time"
)

type User struct {
	ID          string `gorm:"unique"`
	Username    string `gorm:"unique" validate:"required"`
	Email       string `gorm:"unique" validate:"required,email"`
	Password    string `validate:"required,min=8,max=20"`
	Photo       string
	EmailVerify bool
	Role        int       `gorm:"default:0"`
	CreatedAt   time.Time `gorm:"<-:create"` // allow read and create, but don't update
}
