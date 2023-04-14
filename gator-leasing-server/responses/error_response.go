package responses

// A generic error that is returned whens something goes wrong
// swagger:response ErrorResponse
type ErrorResponse struct {
	// The error message
	// in: body
	Body struct {
		// The error message
		// required: true
		Msg string
	}
}
