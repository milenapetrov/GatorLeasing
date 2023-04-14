package responses

// Response returned after creating a new lease
// swagger:response PostLeaseResponse
type PostLeaseResponse struct {
	// in: body
	Body struct {
		// Lease id of created lease
		// required: true
		ID uint
	}
}
