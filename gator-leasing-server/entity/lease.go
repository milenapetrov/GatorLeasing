package entity

import "GatorLeasing/gator-leasing-server/model"

type Lease struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewLease(lease *model.Lease) *Lease {
	return &Lease{
		ID:   lease.ID,
		Name: lease.Name,
	}
}
