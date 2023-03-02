package entity

import "github.com/milenapetrov/GatorLeasing/gator-leasing-server/enums"

type UserContext struct {
	ID        uint
	UserID    string
	InvitedAs enums.InvitedAs
}

func NewUserContext() *UserContext {
	return &UserContext{}
}
