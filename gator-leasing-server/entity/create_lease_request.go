package entity

type CreateLeaseRequest struct {
	Name string `json:"name" validate:"required"`
}
