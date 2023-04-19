package dto

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	OwnerID    uint           `gorm:"primaryKey" faker:"-"`
	OwnerType  string         `gorm:"primaryKey" faker:"-"`
	CreatedAt  time.Time      `faker:"-"`
	UpdatedAt  time.Time      `faker:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" faker:"-"`
	Street     string         `gorm:"not null" faker:"streetFaker"`
	RoomNumber string         `faker:"-"`
	City       string         `gorm:"not null" faker:"cityFaker"`
	State      string         `gorm:"not null" faker:"stateFaker"`
	ZipCode    string         `gorm:"not null" faker:"zipCodeFaker"`
}
