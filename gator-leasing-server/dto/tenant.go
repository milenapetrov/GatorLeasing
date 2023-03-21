package dto

import (
	"time"

	"gorm.io/gorm"
)

type Tenant struct {
	ID        uint `gorm:"primarykey" faker:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}
