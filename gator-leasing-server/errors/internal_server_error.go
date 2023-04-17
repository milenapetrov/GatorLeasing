package errors

type InternalServerError struct {
	Msg string `json:"error"`
}

func (e *InternalServerError) Error() string {
	return e.Msg
}
