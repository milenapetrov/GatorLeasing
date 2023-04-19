package dto

import (
	"time"

	"gorm.io/gorm"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/enums"
)

type TenantUser struct {
	ID          uint            `gorm:"primarykey" faker:"-"`
	CreatedAt   time.Time       `faker:"-"`
	UpdatedAt   time.Time       `faker:"-"`
	DeletedAt   gorm.DeletedAt  `gorm:"index" faker:"-"`
	UserID      string          `faker:"len=30"`
	TenantID    uint            `gorm:"not null" faker:"-"`
	Tenant      Tenant          `faker:"-"`
	InvitedAs   enums.InvitedAs `faker:"-"`
	Email       string          `faker:"email"`
	PhoneNumber string          `faker:"phone_number"`
}
