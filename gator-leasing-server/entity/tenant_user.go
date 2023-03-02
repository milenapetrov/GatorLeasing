package entity

import "github.com/milenapetrov/GatorLeasing/gator-leasing-server/enums"

type TenantUser struct {
	ID        uint
	UserID    string
	TenantID  uint
	InvitedAs enums.InvitedAs
}
