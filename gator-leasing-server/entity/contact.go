package entity

// swagger:model Contact
type Contact struct {
	// the id for this contact
	ID uint
	// the last name for this contact
	LastName string
	// the first name for this contact
	FirstName string
	// the salutation for this contact
	Salutation string
	// the lease id for this contact
	LeaseID uint
	// the phone number for this contact
	PhoneNumber string
	// the email for this contact
	Email string
	// the address for this contact
	Address Address
}
