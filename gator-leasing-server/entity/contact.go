package entity

type Contact struct {
	ID          uint
	LastName    string
	FirstName   string
	Salutation  string
	LeaseID     uint
	PhoneNumber string
	Email       string
	Address     Address
}
