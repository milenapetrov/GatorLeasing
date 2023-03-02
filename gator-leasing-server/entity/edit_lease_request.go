package entity

type EditLeaseRequest struct {
	ID   uint   `json:"id" validate:"required"`
	Name string `json:"name"`
}
