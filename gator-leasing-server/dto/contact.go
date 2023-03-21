package dto

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	ID          uint `gorm:"primarykey" faker:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	LastName    string         `gorm:"type:varchar(30);not null"`
	FirstName   string         `gorm:"type:varchar(30)"`
	Salutation  string         `gorm:"type:varchar(30)"`
	LeaseID     uint           `gorm:"not null"`
	PhoneNumber string         `gorm:"type:varchar(30)"`
	Email       string         `gorm:"type:varchar(30)"`
	Address     Address        `gorm:"polymorphic:Owner"`
}
