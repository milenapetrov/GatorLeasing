package dto

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	OwnerID    uint   `gorm:"primaryKey"`
	OwnerType  string `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Street     string         `gorm:"not null"`
	RoomNumber string
	City       string `gorm:"not null"`
	State      string `gorm:"not null"`
	ZipCode    string `gorm:"not null"`
}
