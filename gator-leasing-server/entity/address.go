package entity

// swagger:model Address
type Address struct {
	// the street for this address
	// required: true
	Street string
	// the room number for this address
	RoomNumber string
	// the city for this address
	// required: true
	City string
	// the state for this address
	// required: true
	State string
	// the zip code for this address
	// required: true
	ZipCode string
}
