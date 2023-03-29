package viewModel

type Address struct {
	Street     string `json:"street" validate:"required"`
	RoomNumber string `json:"roomNumber" validate:"required"`
	City       string `json:"city" validate:"required"`
	State      string `json:"state" validate:"required"`
	ZipCode    string `json:"zipCode" validate:"required"`
}
