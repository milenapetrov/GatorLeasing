package model

import (
	"time"

	"gorm.io/gorm"

	"GatorLeasing/gator-leasing-server/enums"
)

type TenantUser struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    string
	TenantID  uint `gorm:"not null"`
	Tenant    Tenant
	InvitedAs enums.InvitedAs
}
