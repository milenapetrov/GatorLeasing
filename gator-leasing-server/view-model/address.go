package viewModel

type Address struct {
	Street     string `json:"street"`
	RoomNumber string `json:"roomNumber"`
	City       string `json:"city"`
	State      string `json:"state"`
	ZipCode    string `json:"zipCode"`
}
