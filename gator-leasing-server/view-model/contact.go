package viewModel

type Contact struct {
	ID          int     `json:"id"`
	LastName    string  `json:"lastName"`
	FirstName   string  `json:"firstName"`
	Salutation  string  `json:"salutation"`
	LeaseID     int     `json:"leaseID"`
	PhoneNumber string  `json:"phoneNumber"`
	Email       string  `json:"email"`
	Address     Address `json:"address"`
}
