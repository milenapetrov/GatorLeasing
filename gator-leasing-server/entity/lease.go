package entity

type Lease struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	OwnerID uint   `json:"ownerID"`
}
