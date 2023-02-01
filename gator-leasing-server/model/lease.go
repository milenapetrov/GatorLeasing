package model

import "gorm.io/gorm"

type Lease struct {
	gorm.Model
	Name string
}
