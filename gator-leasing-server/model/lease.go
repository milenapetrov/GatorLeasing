package model

import (
	"time"

	"gorm.io/gorm"
)

type Lease struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	OwnerID   uint       `gorm:"not null"`
	Owner     TenantUser `gorm:"foreignKey:OwnerID"`
}
